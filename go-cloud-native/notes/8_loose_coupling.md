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