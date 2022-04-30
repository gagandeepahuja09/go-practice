Stability Patterns
* Improve the service's stability and the larger system that they are part of.

**************************************************************************************

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

**************************************************************************************

Debounce
* Limits the frequency of a function invocation so that only the first or last in a cluster of calls is actually performed.

**Applicability**
* Scenario: When we are performing a series of potentially slow or costly operations where only one would do.
* This technique has been used in JS world for years. Examples:
    * Spam clicking a button to find that all requests after the first request got ignored.
    * Autocomplete pop-up only displays after you stop typing.
* In backend, can be used to retrieve some slowly updating remote resource without bogging down, wasting both client and server time with wasteful resources.
* It is similar to throttling. Difference is that while debounce restricts clusters of invocations, throttle simply limits according to time period.

**Participants**
Circuit: The function to regulate.
Debounce: A closure with the same function signature as circuit.

**Implementation**
* Two approaches:
    * Function-first
    * Function-last

* Function-first:
    * On each call to the function, a time interval is set(regardless of the outcome).
    * Any subsequent  call made before the time interval expires is ignored.
    * We ensure thread safety by using mutex locks. Any overlapping calls will wait for the cached result to be available.
    * Each call will reset the threshold timeout so that call is made exactly once for a series of call.

* Function-last: 
    * More common in FE services like JS.