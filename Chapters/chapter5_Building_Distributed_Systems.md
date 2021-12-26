About the app
* We'll collect and count votes by querying Twitter's streaming API for mentions of specific hash tags.
* Each component will be capable of horizontal scaling to meet demand.

Learnings:
* Interacting with NoSQL datastores.
* Distributed message queue like NSQ.
* Stream live data through twitter's streaming APIs and manage long running connections.
* Stopping programs with internal goroutines.
* Using low memory channels for goroutines.

System Design:
https://docs.google.com/document/d/1xJEB6Sefr7H7kJKOx9Fjv8e-wz1MRX3qX9awzltRBTQ/edit

* Twitter's streaming API allows long-running connections where data is streamed as quickly as possible.
* twittervotes pulls the relevant tweet data via the twitter API, decides what is being push for and pushes the vote in NSQ.
* NSQ makes it available to anyone who has expressed an interest in the vote data.
* counter listens for votes in NSQ and periodically save the results in mongoDB. It keeps an in-memory tally of the result and periodically pushes the result in mongoDB.

* It could be argued that a single program could have done it all. Read the tweet, count the votes and push them to UI. Such a solution would be very limited in scale.
* In our solution any of the components can be scaled as the demand of that particular capability increases.
* If we have relatively fewer polls but lots of people viewing the data, we can keep the twittervotes and counter instances down and add more web and mongoDB instances or vice-versa if the situation is reversed.
* Another key advantage is redundancy and avoiding SPOF. If one of our boxes disappear, then others can pick up the slack.


Database Design
* We'll create a mongoDB database called ballots. It will contain a single collection polls.
* fields: title, options(array), result(json)

Installing the environment
* Services such as nsqd and mongod will have to be started before we can run our programs.
* NSQ guarantees atleast once delivery. It means that it keeps undelivered messages cached until all interested parties have received them.
* This means that even if we stop our counter program, we won't miss any votes.
* The key tools from NSQ installation that we are going to use are: nsqlookupd, nsqd.
* nsqlookupd is a deamon that manages topology information about the distributed NSQ environment. It keeps track of all the nsqd producers for specific topics and provides interfaces for clients to query such information.
* nsqd is the deamon that does the heavy lifting for NSQ, such as receiving, queuing and delivering message to and from interested parties.
* go get github.com/nsqio/go-nsq

Installing MongoDB
* brew services tap and brew install command.
    To start mongodb/brew/mongodb-community now and restart at login:
    brew services start mongodb/brew/mongodb-community
    Or, if you don't want/need a background service you can just run:
    mongod --config /usr/local/etc/mongod.conf

MongoDB driver for go
* go get gopkg.in/mgo.v2

Starting the environment
* Start nsqlookupd so that our nsqd environments are discoverable.
* Start nsqd and tell which nsqlookup to use.
* Start mongod for data services.

1. nsqlookupd
2. nsqd --lookupd-tcp-address=localhost:4160
3. sudo mongod --dbpath /usr/local/var/mongodb/

* All three in separate terminals.
* dbpath is where the data will be stored


Reading Votes From Twitter(twittervotes)

twitter votes is going to do the following:
* Load all polls from the mongoDB database using mgo and collect all options from options array in each document.
* Open and maintain a connection to Twitter's streaming APIs looking for any mention of the options.
* Figure out which option is mentioned and push that option through to NSQ for each tweet that matches the filter.
* If the connection to twitter is dropped(which is common in long-running connections that are part of Twitter's streaming API connection), then after a short delay(so that we don't bombard Twitter with connection requests), reconnect and continue.
* Periodically re-query MongoDB for the latest polls and refresh the connection to Twitter to make sure that we are always looking out for the right options.
* Gracefully stop itself when the user terminates the connection by pressing Ctrl+C.
