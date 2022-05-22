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