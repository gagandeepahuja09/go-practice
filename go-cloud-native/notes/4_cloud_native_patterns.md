Stability Patterns
* Improve the service's stability and the larger system that they are part of.

Circuit Breaker
* Degrades service functions in response to a likely fault.
* Temporarily ceases to execute requests.
* Prevents larger/cascading failures by eliminating recurring errors and providing reasonable error responses.

**Applicability**
* Service failures inevitable: misconfigurations, database crashes, network partitions, etc.
* If not done can lead to:
    * Wastage of resources.
    * Obscuring the source of original failure.
    * Cascading failures. 

**Participants**
* Circuit: The function that interacts with the service.
* Breaker: A closure with the same function signature as circuit.

**Implementation**
* The circuit breaker is just a specialized Adapter pattern.
* Breaker wraps circuit to add some additional error handling logic.
* Breaker takes circuit as input parameter and returns a circuit.
* We'll later review why exponential backoff isn't the ideal algorithm.