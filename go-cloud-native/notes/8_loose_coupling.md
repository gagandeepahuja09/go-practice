* Coupling ==> Loose v/s Tight coupling.
* Section 1: 
    * Some kinds of tight coupling can lead to the dreaded "distributed monolith".
* Section 2: 
    * Inter-service communications.
    * Fragile exchange protocols can lead to tight coupling.
* Section 3: 
    * Implementation of the services.
    * Use of plug-ins as a way to dynamically add implementations.
* Section 4:
    * Hexagonal architecture: An architectural pattern that makes loose coupling the central pillar of its design philosophy.

**Tight Coupling**
* Coupling describes the degree of direct knowledge between components.
* A client that sends a request to a service is by definition coupled to that service.
* Tightly coupled components have a great deal of knowledge about another component. Eg:
    * Both could require the same version of a shared library to communicate.
    * Client could require an understanding of the server's architecture or DB schema.
* It's easy to build tightly coupled systems when optimizing for the short term, but they have a huge downside: the more tightly coupled two components are, the more likely that a change to one component will necessitate corresponding changes to the other. As a result, tightly coupled systems loose many of the benefits of a microservice architecture.

**Loose Coupling**
* Loosely coupled components have minimal direct knowledge of one another. They are relatively independent.
* Systems designed for loose coupling require more up-front planning, but they can be more freely upgraded, redeployed, rewritten without greatly affecting the system that depend on them.

* If you want to know how tightly coupled your system is, ask how many and what kind of changes can be made to one component without adversely affecting another.

**Coupling in Different Computing Contexts**
* In programming, code can be tightly coupled when a dependent class directly references a concrete implementation instead of an abstraction(e.g interface).
* In Go, this might be a function that requires an os.File when a io.Reader would do.
* Multiprocessor systems that communicate by sharing memory can be said to be tightly coupled. In a loosely coupled system, components are connected through a MTS(Message Transfer System) like Channels.
* Occasionally, it might be good to tightly couple certain components. Eliminating abstractions and other immediate layers can reduce overhead which can be a useful optimization if speed is a critical system requirement.

**Tight coupling takes many forms**
* They all share one fundamental flaw - they all depend on some property of another component that they wrongly assumed won't change.
* They can be divided into different classes depending on the resource they are coupled to.

**Fragile Exchange Protocols**
* SOAP(Simple Object Access Protocol) service provided a contract that clients could follow to format their requests(in XML).
    * It was very fragile: if the contract changed in any way, the clients had to be updated along with it.
    * Hence we can say that SOAP clients were tightly coupled to their services.
* It was quickly replaced by REST which was a considerable improvement but can introduce its own tight coupling.
* gRPC included a number of important features, including allowing loose coupling between components.

**Shared dependencies: Distributed Monolith**
* We all know that microservice architecture is generally more desirable than monolith but that is easier said than done.
* This is because, it's so easy to accidentally create a distributed monolith: a microservice-based system containing tightly coupled components.
* Distributed Monolith ==> Worst of all worlds.
* Problems:
    * Services often can't be deployed independently, so deployments have to be carefully orchestrated.
    * Cascading failures.
    * Rollbacks are functionally impossible.

**Shared point-in-time**
* Request-response messaging pattern - clients expect an immediate response from services. It it's not, the services will fail. It can be said that they are coupled in time.
* Coupling in time is not necessarily a bad thing and at times even preferable.
* But, if the response isn't time constrained, then a safer approach may be to send messages to an intermediate queue.

**Fixed Address**
* It's the nature of microservices that they need to talk to one another. But for that, they need to find each other.
* *Service discovery*: Process of locating services on a network.

* Traditionally this is discovered from some centralized registry. Initally it was in the form of manually maintained files like hosts.txt.
* As networks scaled, DNS started getting adopted.
* DNS works well for long-lived services whose location on the network rarely change.
* With ephemeral, microservices-based applications, the service lifespan is reduced to minutes/seconds rather than years/months.
* In such dynamic environments, URLs and traditional DNS has become another form of tight coupling.
* This need for dynamic, fluid service discovery has driven the adoption of entirely new strategies like the *service mesh*, a dedicated layer for facilitating service-to-service communications between resources in a distributed system.

*******************************************************************************

**Communication Between Services**
* In order for services to communicate, they must first establish an implicit or explicit that defines how messages will be structured.
* While such a contract is necessary, it also effectively couples the components that depend on it.

* Does the protocol allow backward- and forward-compatible changes, like protocol buffers and gRPC?
* Or do minor changes to the contract effectively break communications, as is the case with SOAP?

**Request-Response Messaging**
* Point-to-point: one request, one receiver.
* Requires the requesting process to pause until it receives a response.

**Common Request-Response Implementations**

* *REST*
    * Human-readable and easy to implement, making it a good choice for outward-facing services.

* *RPC*
    * RCP frameworks allow programs to execute procedures in a different address space, often on another computer.
    * 2 big language agnostic RPC players: gRPC, Apache Thrift.

* *GraphQL*
    * Query & manipulation language generally considered an alternative to REST.
    * Particularly powerful when working with complex datasets.


**Issuing HTTP Requests with net/http**

* net/http includes convenience functions for GET, HEAD, POST methods.
func Get(url string) (*http.Response, error)

* A small selection of the http.Response struct:

type Response struct {
    Status          string  // eg. "200 OK"
    StatusCode      int     // eg. 200

    // Header map header keys to values
    Header Header

    // Body represents the response body
    Body io.ReadCloser

    // ContentLength records the length of the associated content. The value
    // -1 indicates that the length is unknown.
    ContentLength int64

    // Request is the request that was sent to obtain this response
    Request *Request
}

* Body field provides access to the HTTP response body.
* It's a ReadCloser interface, which tells us 2 things:
    * The response body is streamed on demand, as it is read.
    * It has a Close method which we are expected to call.
* Failing to close our body can lead to memory leaks.
* io.ReadAll of response.Body would return a []byte slice.

**Issuing POST HTTP Requests**

* There are 2 convenience functions: Post, PostForm.

// Post issues a POST to the specified URL.
func Post(url, contentType string, body io.Reader) (*Response, error)

// PostForm issues a POST to the specified URL, with data's keys 
// and values URL-encoded as the request body.
func PostForm(url string, data url.Values) (*Response, error) 

**A Possible Pitfall of Convenience functions**
* What does convenience functions mean?
* If we check the implementation of http.Get, we can see that it calls DefaultClient.Get.
* Because http.Client is concurrency safe, it's possible to have only one of these pre-defined as a package variable.
* DefaultClient is a zero-value http.Client.
    var DefaultClient = &Client{}
* This has a timeout of 0. Which is interpreted as no timeout.
* If the server doesn't respond and doesn't close the connection, then it could result in a nondeterministic memory leak.
* For adding timeouts, we need to add a custom client.

var client = &http.Client{
    Timeout: time.Second * 10,
}
response, err := client.Get(url)

**Remote Procedure Calls With gRPC**
* Efficient, polyglot data exchange framework.
* REST is essentially a set of unenforced best practices.
* RPC frameworks allow a client to execute a specific method implemented on different systems as if they were local functions.

Advantages:

* *Conciseness*: Its messages are more compact, consuming less network I/O.

* *Speed*: Its binary exchange format is much faster to marshal and unmarshal.

* *Strong-typing*: Eliminates a lot of boilerplate and removes common sources of errors.

* *Feature-rich*: No. of built-in features like authentication, encryption, compression, timeouts.

Few Possible Disadvantages:

* *Contract driven: Only internal services*: gRPC's contracts make it less suitable for external facing services.

* *Binary format: Not human readable*: gRPC data isn't human readable, making it harder to inspect and debug.

**Interface Definition With Protocol Buffers(IDL)**
* As with most RPC frameworks, gRPC requires you to define a service interface.
* gRPC frameworks implement the resulting source code to handle client calls.
* The client has a stub that provides the same methods as the server.

**The message definition structure**
* *Protobufs*: 
    * Language-neutral mechanism for serializing structured data.
    * Kind of like binary version of XML.
* Each field in a message definition has a field number that is unique for that message type. 
    * They are used to identify fields in the message binary format, and should not be changed once your message type is in use.
    * Protobufs ensure that it doesn't lead to tight coupling:
        * Fields can be removed, as long as the field number is not used again in your updated message type.
        * We can mark the field as reserved so that they can't be accidentally reused.
* C++ like syntax.

**The key-value message structure** 
* Implementing gRPC equivalents to the Get, Put and Delete functions we are already exposing via RESTful methods.
* We don't include error and status in the message response definitions. They are unnecessary because they are included in the return values of the gRPC client functions.

**Defining our service methods**

**Compiling protocol buffers**
* Using protoc command.
* We use the --go_out option since we want go code.
* go_opt and go-grpc_opt tells protoc to place the output files in the same relative directory as the input file.
* 2 files created ==> keyvalue.pb.go and keyvalue_grpc.pb.go
* 
    protoc --proto_path=$SOURCE_DIR \
    --go_out=$DEST_DIR --go_opt=paths=source_relative \
    --go-grpc_out=$DEST_DIR --go-grpc_opt=paths=source_relative \
    $SOURCE_DIR/keyvalue.proto
* go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest                               
* export SOURCE_DIR=go-cloud-native/distributed-key-val-store/keyvalue
    if you are in go-practive directory.

**Implementing the gRPC service**
* To implement our gRPC server, we'll need to implement the generated service interface. Can be found in keyvalue_grpc.pb.go.
* In KeyValueServer interface, each method accepts a context.Context and a request pointer and returns a response pointer and an error.
* As a side effect of its simplicity(interfaces), it becomes quite easy to mock requests to, and responses from a gRPC server implementation. 
* We can implement our key value gRPC server by embedding the UnimplementedKeyValueServer struct.
* The UnimplementedKeyValueServer by default already implements all the methods.
* Steps involved:
    1. *Create the server struct*: It should embed UnimplementedXXXServer
    2. *Implement the service methods*: Since UnimplementedXXXServer includes for all the service methods, we don't have to implement them right away.
    3. *Register our gRPC server*: Create a new instance of server struct and register it with the gRPC framework.
    4. *Start accepting connections*.
* gRPC provides the best of both worlds by providing the freedom to implement any of the desired functionality without having to be concerned with building many of the tests and checks usually associated with a RESTful service.

**Implementing the gRPC Client**
* The generate client interface will be name XXXClient

    type KeyValueClient interface {
        Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error)
        ...
    }

* All of the client methods(stubs) are already implemented, hence we need not worry about their implementation.
* Each method accepts a request type pointer, and returns a response type pointer and an error.
* Additionally each method accepts a context.Context and zero or more instances of grpc.CallOption.
* CallOption is used to modify the behaviour of the client when it executes its calls.
* Options used in our example:
    * WithInsecure: Disables transport security for its ClientConn. Don't use insecure connections in production.
    * WithBlock: Makes dial a block until a connection is established, otherwise the collection will occur in the background.
    * WithTimeout: Makes a blocking dial throw an error if it takes longer than specified amount of time.
* **Best Part**: These functions look and feel exactly like function calls. No checking status-codes, hand-building our own clients, or any other funny business.

**********************************************************************************

**Loose Coupling Local Resources with Plug-ins**
* It's often useful to build services or tools that can accept data from different kinds of input sources(such as REST interface, a gRPC interface, and a chatbot interface) or generate different kinds of outputs(such as generating different kinds of logging or metric formats).
* Added bonus: designs that support such modularity can also make mocking resources for testing dead simple. 

**In-Process Plug-ins with the plugin Package**
* Go provides native plug-in support via standard *plugin package*.
    * It's used to open and access go plug-ins. Not necessary to actually build the plug-in themselves.
* 3 requirements of Go plug-in:
    * It must be in the main package.
    * It must export one or more functions or variables.
    * It must be compiled using the -buildmode=plugin build flag.

**********************************************************************************

**Hexagonal Architecture**
* aka "Ports and Adapters" pattern.
* Uses *loose coupling* and *inversion of control* as its central design philosophy to establish clear boundaries b/w business and peripheral logic.
* In a hexagonal application, the core application doesn't know any details at all about the outside world, operating entirely through loosely couple *ports* and technology specific *adapters*. 

* This approach allows the application to expose different APIs(REST, gRPC, a test harness, etc) or use different data sources(database, message queues, local files, etc) without impacting its core logic or requiring major code changes.