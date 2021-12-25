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