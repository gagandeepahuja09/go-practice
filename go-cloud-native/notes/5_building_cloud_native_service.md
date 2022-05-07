* Building a distributed key value store.

Requirements 
* Store key-value pairs.
* Service endpoints.
* Persistence.
* Idempotence

* Idempotence: Calling a function once has the same effect as calling it multiple times.
* Nullipotence: Trigger no state change. Eg. x = 1, HTTP PUT is idempotent but not nullipotent. HTTP GET, x=x are nullipotent as well.

Why do we need idempotence?
* Safer: Retries would give unexepected results in not idempotent.
* Idempotent operations are often simpler: Eg. a PUT method requiring only setting value vs. a POST method that checks whether a value is set and depending on that throwing an error.
* Idempotent operations are declarative 
    * They tell a service what needs to be done(declarative), instead of telling how to do it(imperative).
    * This frees users from dealing with low-level constructs and minimizes potential side-effects.
* Cloud-native services must hold this property.

************************************************************************************

Generation 0: The Core Functionality
* While returning the error, we don't use errors.New. Instead we return the prebuilt ErrNoSuchKey error value.
* This is an example of *sentinel error*, which determines exactly what type of error it's receiving and to respond accordingly.

    if err.Is(ErrNoSuchKey) {
        http.Error(w, err.Error(), http.StatusNotFound)
        return
    }

***********************************************************************************

**Generation 1: The Monolith**
* REST => Perfectly sufficient for our needs.
* Go standard libraries are designed to be extensible, so there are a no. of web frameworks that extend them. 
* Handler => any type that implements the Handler interface.
* mux/multiplexer => direct incoming signals to one of the possible outputs.
* When a request is received by a service that's been started by ListenAndServe, it's the job of a mux to compare the request URL to the registered patterns and call the handler function with the one that matches closely.
* If mux is specified as nil, the DefaultServeMux is used.

gorilla/mux
* For many of the use-cases, DefaultServeMux should be fine.
* One of the key features of gorilla/mux is the ability to create paths with variable segments which can optionally contain a regular expression pattern.
* Format => {name} or {name:pattern}
    r := mux.NewRouter()
    r.HandleFunc("/products/{key}", ProductHandler)
    r.HandleFunc("/articles/{category}/", ArticlesCategoryHandler)
    r.HandleFunc("/articles/{category}/{id:[0-9]+}, ArticleHandler)
* We can retrieve the variables using mux.Vars function which returns a map[string]string.
* Matchers => We can add additional matching criterias. Eg domains/subdomains, methods, schemes, path prefixes, headers or even custom matching functions that we create. 
    * Each matcher returns a route, hence can be easily chained.
    * r.HandleFunc("/products", ProductHandler).
        Host("www.example.com").
        Methods("GET", "PUT").
        Schemes("http")

Building a RESTful service
* Built the GET, PUT methods, both of which are idempotent.

**Making Our Data Structure Concurrency Safe**
* Maps in Go are not atomic and not safe for concurrent use.
* Using magic of composition, we'll create an *anonymous struct* that contains our map and an embedded sync.RWMutex.

***********************************************************************************

**Generation 2: Persisting Resource State**
* One of the stickiest challenges with distribute cloud native applications is how to handle state. 
* 2 Common ways to maintain the state of the application:
    * Transaction Log File: Makes sense when we want to store our data in memory most of the time, only accessing persistence mechanism during startup time.
    * External Database: Will help with scaling across multiple replicas and provide resilience. 

**Application State v/s Resource State**
Application State:
    * Server-side data about the application or how it's being used by a client.
    * Eg. client session tracking, eg to associate them with their access credentials or some other application context.
Resource State:
    * The current state of a resource within a service at any point in time.
    * It's the same for every client and has nothing to do with the interaction b/w client and server.

* Application state can be more problematic because it requires *server-affinity*. Which means sending each of a user's request to the same server(stickiness). This can make it hard to destroy/replace server replicas.

**What's A Transaction Log?**
* Assuming that it'll be read only when the service is restarted or otherwise needs to recover its state.
* It'll be read from top to bottom, sequentially replaying each event.
* For speed & simplicity ==> append only.
* Attributes:
    * Sequence Number
    * Event Type(PUT/DELETE)
    * Key
    * Value
* Create a txn logger interface with 2 methods for now: WritePut, WriteDelete.

**Storing State In A Transaction Log File**
Pros:
    * No downstream dependency on any external service.
    * Technically straightforward.
Cons:
    * Harder to scale: Will need additional logic to distribute our state between nodes when we want to scale.
    * Uncontrolled Growth: We'll need to find a way to compact them(something like LSM trees).

Prototyping: 
* We'll write in plain-text. A binary, compressed format might be more time and space efficient, but we can always optimize later.
* Each entry will be written in its own line. (easier to read).
* Each line will include the 4 fields delimited by tabs.

Defining Event Type
* We want the WritePut and WriteDelete methods to operate asynchronously.

Declaring Constants With Iota

* It's value starts with zero in each constant declaration and increments with each constant assignment, whether or not the iota identifier is actually referenced.

const (
    a = 42 * iota   // iota = 0, a = 0
    b = 1 << iota   // iota = 1, b = 1
    c = 3           // iota = 2, c = 3  (iota increments anyway!)
    d float64 = iota / 2 // ioat = 3, d = 1.5 
)

* The iota keyword allows *implicit repetition*, which makes trivial to create arbitrarily long sets of related constants as show in below bytes example.

type ByteSize uint64

const (
    _ = iota
    KB ByteSize = 1 << (10 * iota) // KB = 2 ^ 10
    MB // 2 ^ (10 * 2)
    GB // 2 ^ (10 * 3)
    TB
    PB
)

**Implementing Your FileTransactionLogger**
* In order to keep a track of the physical location of the transaction log, we'll have an os.File attribute.
* Keeping an io.Writer that the WritePut and WriteDelete methods will operate on directly would be a single-threaded approach. We might be spending a lot of time in I/O.
* We'll instead write events to a channel which will be processed by a goroutine running in parallel.
* While the above approach makes for more efficient writes, it means that WritePut, WriteDelete can't return an error. We'll use a dedicated errors channel to deal with that instead.

**Creating A new FileTransactionLogger**
* NewFileTransactionLogger will call the os.OpenFile function to open the file specified by the filename parameter.
* Several flags ORed to set its behavior:
    * os.O_RDWR: Opens the file in read/write mode.
    * os.O_APPEND: Any writes to the file will be append, not overwrite.
    * os.O_CREATE: If the file doesn't exist, create it.
* We could spawn our goroutine for listening to the events channel in this NewFileTransactionLogger method but that would make it look more mysterious.
* Instead we should have a separate Run method.

**Appending entries to the transaction log file**
* Using a buffered channel ensures that the call to WriteDelete and WritePut won't fail as long as the buffer isn't full.
* This lets the consuming event handle short bursts of events without being slowed by disk I/O.
* If the buffer does fill up, then the write methods will block until the log writing goroutine catches up.
* The errors channel is also buffered with size of 1 to ensure that we are able to send the error in a non-blocking manner. We halt the goroutine when we get an error as it has a size of only 1.

**Using a bufio.Scanner to play back file transaction logs**
* We'll use fmt.Sscanf to read through it and extract the four necessary attributes.
* Go's concurrency primitives makes it trivially easy to stream the data to the consumer in a much more space and memory-efficient way.
* The ReadEvents method can be said to be two functions in one:
    * Outer function initializes the bufio.Scanner and returns the event and error channels.
    * The inner function run concurrently to ingest the file contents line-by-line and send the results to the channels.
* file attribute in transaction logger is of *os.File which has a read method.
* We reuse the same Event value in each iteration rather than creating a new one. This is because outEvent channel is sending struct values and not pointer to struct. 

**Updating the transaction logger interface**

**Initializing the FileTransactionLogger in our web service**
* foo, ok = <-ch; ok will be false if the channel is closed.
* Alternatively we can range through the channel. The channel closes when the range ends.

**Integrating the FileTransactionLogger in our web service**

**Future Improvements**
* No tests.
* No close method to gracefully close the file.
* The service can close with events still in write buffer; events can get lost.
* Keys and values aren't encoded in the transaction log: multiple line or white space will fail to parse correctly.
* Unbound size of keys and values.
* Writing in plain-text will take up more space.

***********************************************************************************

**Storing State In An External Database**
* Part of standard go library: https://pkg.go.dev/database/sql

Pros:
* Externalizes application state
    Less need to worry about distribute state and closer to "cloud native".
* Easier to scale
    Not having to share data between replicas makes scaling out easier(not easy).
Cons:
* Introduces an upstream dependency
    Creates a dependency on another resource that might fail.
* Increases complexity
    Yet another thing to manage and configure.                        

**Working with databases in Go**
* databases/sql package provided by Go.
* (sql.DB) => most common member of the package.
    * Go's primary DB abstraction.
    * Entry point for creating statements and transactions, executing queries, fetching results.
    * Negotiates connection with the DB and maintains a connection pool.

**Importing A database driver**
* sql.DB => common interface for interacting with a SQL DB.
* Database driver => implements the specifics of a DB type.
* The lib/pq postgres driver package will be imported anonymously.

**Defining PostgresTransactionLogger struct**

**Creating a new PostgresTransactionLogger**

**Using db.Exec to execute a SQL Insert**

**Using db.Query to play back transaction logs**
