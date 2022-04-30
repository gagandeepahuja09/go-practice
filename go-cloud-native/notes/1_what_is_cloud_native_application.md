Major components: Scalable, Loosely Coupled, Resilient, Manageable, Observable.

Resilience
* Roughly synonymous with fault tolerance.
* How well a system withstands and recovers from errors and faults.
* A system can be considered resilient if it can continue operating - possibly at a reduced level - rather than failing completely when some part of the system fails.
* You can't prevent every possible fault and it's wasteful and unproductive to try.
* By assuming that all of system's components are certain to fail - which they are - and designing them to respond to potential faults and limit the effects of failures, we can produce a system that's functionally healthy even when some of its components are not.
* Many ways:
    * Deploying redundant components. (A fault shouldn't affect all components of the same type).
    * Circuit breakers and retry logic.
* Resilience is not reliability:
    * Resilience: degree to which a system can continue to operate