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


Authorization with Twitter
* Step1: Create an app in developer console of Twitter.
* Create a new file called setup.sh to keep the creds.


Extracting the Connection
* In order to maintain long running connections, we'll create our very own dial method. This method will close the existing connection if it exists and setup a new connection with a timeout of 5 seconds. This connection will be maintained in a global variable conn.
* If a connection dies or is closed by us, we can safely call this without having to worry about zombie connections.
* We will periodically close the connection ourselves and initiate a new one because we want to reload the option from DB at regular intervals.
* We'll also close io.ReadCloser in this method which we'll use to read the body of the responses.

Reading Environment Variables And Setting up Auth
* go get github.com/gomodule/oauth1/oauth
* go get github.com/joeshaw/envdecode

Making Request To Twitter
* Singleton Pattern: We use sync.once to ensure that our initialization code runs only once despite the number of times we call makeRequest.
* After calling setupTwitterAuth, we create a new http.Client function using a http.Transport function that uses our custom dial method. 
* sync.Once.Do majorly does mutex lock and defers the unlock.
* For making the request, we'll use an http client. 
* In order to describe the mechanism by which the request is to be made, we'll use http.Transport.
* Dial method in http.Transport is deprecated. If both are present, then dialContext takes prioriy. 
* dialContext allows to cancel dials as soon as they are no longer needed.
* We'll also set important headers like Authorization headers, content-type, content-length.
* for creating the authorization header string, we'll use oauth.Client.AuthorizationHeader method.
* One limitation is that we have used oAuth 1 in code.
* In the end we simply return the response of the call to client.Do. 

Reading from MongoDB
* dialdb, closedb functions will connect to and disconnect from the locally running mongodb instance.
* It'll store mgo.Session(the database connection object) in a global variable called db. 

Load Options From DB
* We need to extract all options from the documents.
* Our poll document contains more than just options but for new we only care about Options so only that will be present in our poll struct.
* Variable of type poll is passed in iter.Next as address so that the struct gets changed accordingly.
* We get an iterator by calling the Iter method which allows us to access each poll one by one.
* This is a very memory efficient way of reading data because it only ever uses one single poll object.
* If we were to use All method instead, the amount of memory used would depend on no. of DBs which could be out of our control.

Scale Issues:
* With millions of polls in the DB, the options slice would be too large. 
* For that kind of scale, we can run multiple twittervotes programs, each dedicated to a portion of poll data.
* A simple way could be to break on the basis of starting character, ie A-N and O-Z. This is similar to sharding but not actually for a particular DB(not DB sharding).
* A more sophisticated approach would be to add a field to the poll document, one the basis of which we'll group the documents in a more balanced way.


Reading from Twitter:
* Now we'll write code that initiates the connection and continuously reads from the stream, until we either call our close connection method or twitter close the connection for some reason.
* The structure of a tweet can be very complex but since we are only concerned with the tweet text, we'll only keep this in our struct.
* This may look incomplete, but it makes our intentions very clear to other programmers who might see our code.
* readFromTwitter method will keep on sending votes to the send only channel, votes. It is send only since it is of chan<- type. This makes our intentions very clear that we never intend to read the votes channel in our readFromTwitter method. 

* Steps
    * Get all options.
    * Using url.Parse, get the appropriate url.URL object.
    * We create a url.Values object called query and set the track field which has all the options.
    * Create a request object using http.CreateRequest.
    * Pass all this to the makeRequest method.
    * We make a new decoder using response body.
    * We keep reading inside an infinite for loop by calling the Decode method.
    * If the tweet has mentioned some options, then send it to the votes channel.


Starting Twitter Stream
* Terminating the program whenever Decode returns an error doesn't provide a very robust solution. Doc states that API will drop the connection from time to time.
* Also we are going to terminate the connection periodically, so we need to think about a way of reconnecting once the connection is dropped.
* Continually calling the readFromTwitter after some delays, eg. 10 second delays will help.
* In order to acheive this we'll have an infinite for loop in our goroutine and after calling the function readFromTwitter, we'll have a 10s delay.
* Signal Channels: To signal events b/w codes running in different goroutines.
* startTwitterStream will start a goroutine that continually calls readFromTwitter until we signal that we want it to stop.
* Once it has stopped, we want to be notified through another signal channel.
* The return of the function will be a struct{}: a signal channel.
* struct{}{}(0 bytes) is a more memory efficient way of signaling events than bool(1 byte)
* fmt.Println(reflect.TypeOf(true).Size())
* fmt.Println(reflect.TypeOf(struct{}{}).Size())
* We are going to use 2 signal channels here. One is for deciding whether it should stop and another to signal the stop so that can be consumed by other goroutines.
* We'll return a read-only channel so that other goroutines can only read from it.
* Outside the goroutine, other goroutines will use stopchan to tell our goroutine to stop. It is receive only inside this function. At other places it will be capable of sending.
* Returning a channel in this function is necessary because our function triggers it's own goroutine and immediately returns, so without returning this calling code would have no idea whether the spawned code was still running or not.
* Time.sleep is to give the twitter API some rest in case it closed the connection due to overuse. Once we've rested, we re-enter the loop and check if calling code wants us to stop or not.
* NOTE: Signal channels are a great way to start and stop something when all code lives in a single package. If we need to cross API boundaries, context package is the recommended way to deal with deadlines, cancellation and stopping.


Publishing to NSQ
The votes channel is received by the publishVotes method, in order to publish these votes to NSQ under the votes topic. This is also achieved by goroutines, like a slice / array, we can iterate towards a channel using ⇒ for vote := range votes {}. After this we stop the publisher and send the stopChan signal. 
* This will take channel of votes as receive only(<-chan) in input.
* The channel we create later will be made using make(chan string) and won't be either receive / send only while declaring.
* By closing the votes channel, the external code tells our function to stop working.
* Once the votes channel is closed, we will stop publishing and send a signal down the returned stop signal channel. 
* When it comes to sending the empty channel, we could have used defer as well.

Gracefully Starting And Stopping Programs
* When our program is terminated, we want to do a few things before exiting, namely closing our connection to Twitter and stopping the NSQ publisher.
* To achieve this, we have to override the default Ctrl+C behaviour.
* We have a bool stop indicating the current status. Since this will be changed by multiple goroutines, we will have an associate sync.Mutex lock.
* We use signal.Notify to ask Go to send the signal down signalChan when someone tries to halt the program.
* <-signalChan ==> by doing this, we are blocking waiting for the signal by trying to read from signalChan.


Finally Calling All Goroutines in Main Method
* We block twitterStoppedChan by attempting to read from it.

