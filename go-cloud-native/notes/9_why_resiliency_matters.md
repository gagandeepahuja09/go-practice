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