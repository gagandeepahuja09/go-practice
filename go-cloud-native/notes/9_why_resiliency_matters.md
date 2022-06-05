* Failures in complex systems generally never have a single root cause.

**Resilience**
    * The ability of a system to *withstand* and *recover* from errors and failures.
    * A system can be considered resilient if it can continue operating correctly - possibly at a reduced level - rather than failing completely when one of its subsystem fails.

**Resilience is not reliability**
* *Resilience*: 
    * Degree to which a system can continue to operate correctly in the face of errors and faults.
    * Resilience along with the other 4 cloud native properties is just one factor that contributes to reliability.
* *Reliability*:
    * Ability of a system to behave as expected for a given time interval.
    * Reliability along with availability and maintainability contributes to a system's overall dependability.

**What does it mean for a system to fail**
* A system is a set of components that work together to accomplish an overall goal.
* Each component of a system - a *subsystem* is also a complete system in itself, that in turn is composed to smaller subsystems and so on and so on.

**Building for Resilience**
* *All components are destined to fail eventually*.
* Designing them to *respond gracefully to errors*, when they do occur, can produce a system that's functionally healthy even when some of its components are not.

* Multiple ways of increasing the resilience of a system:
    * *Redundancy*: Deploying multiple components of the same type.
    * Specialize logic like *circuit breakers* and *request throttles* to isolate specific kinds of errors and preventing them from propagating.
    * Fault components can be *reaped(cut)* or *intentionally allowed to fail* to benefit the health of the larger system.

**Cascading failure**
* One part of a system experiences a local failure - a reduction in capacity, an increase in latency, etc.
* This causes the other components to attempt to compensate for the failed component in a way that exacarbates the problem, eventually leading to the failure of the entire system.

* The classis cause of cascading failures is **overload**.
    * Occurs when one or more node in a set fails, causing the load to be catastrophically redistributed to its survivors.
    * The increase in load overloads the remaining nodes, causing them to fail from resource exhaustion, taking the entire system down.
    * It can become difficult to scale our way out of the problem. New nodes can be overloaded as quickly as they come online, often contributing the feedback to the system down in the first place.
    * Sometimes the only fix is to take your entire service down - by explicitly blocking the problematic traffic and then slowly reintroduce load.

**Preventing Overload**
* For *every service* there exists some *request frequency*, a threshold beyond which bad things will start to happen.

Strategies

**Throttling**
* Limits the no. of requests that a user can make in a certain period of time.
* Throttles are generally applied on a per-user basis to provide something like a usage quota, so that no one user can consume too much of service's resources.
* We'll need a separate bucket for each user as per the token bucket implementation.
* Throttle doesn't return an error when it's activated: it isn't an error, so we don't treat it as one.
* The current implementation won't use a time.Ticker to explicitly add tokens to buckets on some regular cadence.
* It refills buckets on demand, based on time elasped between requests. This strategy means that we don't have to dedicate background processes to filling buckets until they are actually used, which will scale much more effectively.

**Load Shedding**

**Graceful service degradation or fallbacks**
* Strategically reducing the amount of work needed to satisfy each request instead of just rejecting requests.
* Common approaches - falling back on cached data or less expensive - if less precise - algorithms.