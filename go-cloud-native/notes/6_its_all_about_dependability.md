* The most important property of a program is whether it accomplishes the intention of its user.(users not creators).

What's the point of cloud native?

**It's all about dependability**
* Every single cloud native pattern exists to allow services to be deployed, operated and maintained at scale in unreliable environments, driven by the need to produce dependable services that keep users happy.

**What is dependability and why it's so important**
* A dependable system consistently does what its users expect and can be quickly fixed when it doesn't.
* It's hard to objectively gauge "user expectations".
* Hence dependability is an umbrella concept encompassing several more specific and quantifiable attributes - availability, reliability, and maintainability - all of which are subject to similar threats that may be overcome by similar means.

* *Availability*: 
    * The ability of a system to perform its intended function at a random moment in time. 
    * Expressed as: Probability that a request made to the system will be successful => Uptime / Total time.

* *Reliability*:
    * The ability of a system to perform its intended function for a given time interval.
    * Often expressed as MTBF: Mean time between failure = total time / no. of failures.
    * Or as failure rate: no. of failures / total time.

* *Maintainability*:
    * The ability of a system to undergo modifications and repairs.
    * Can be measured by tracking the amount of time required to change a system's behavior to meet new requirements or to restore it to a functional state. 

**Dependability: It's not just for ops anymore**
* On operations side: With the availability of infrastructure and platforms as a service(IaaS/PaaS) and tools like Terraform & Ansible, working with infrastructure has never been more like writing software.
* On dev side: popularization of technologies like containers and serverless functions has given devs an entire set of "operations-like" capabilities particularly around virtualization and deployment.
* As a result, the once stark line b/w software & infra is getting increasingly blurry. One could argue that everything is software now.

**Acheiving dependability**
Fault Prevention: Used during system construction to prevent the occurence or introduction of faults.
Fault Tolerance: Used during system design & implementation to prevent service failures in the presence of faults.
Fault Removal: Fault removal techniques to reduce the number and severity of faults.
Fault Forecasting: Identify the presence, creation, and consequence of faults.

* These 4 means correspond very well to the 5 cloud native attributes.

********************************************************************************************************

**Fault Prevention**
* Many - if not most - classes of errors and faults can be predicted and prevented during the earliest phases of development.

**Good programming practices**
* Explicit goal of any development methodology is fault prevention - from pair programming, to TDD, to CR practices.

**Language Features**
* Features such as dynamic typing, pointer arithmetic, manual memory management, and thrown exceptions can easily introduce unintended behaviors that are difficult to find and fix, and may even be maliciously exploitable.

**Scalability**
* Ability of a system to continue to provide correct service in the face of significant change in demand.
* We'll discuss horizontal and vertical scaling in detail later.
* We'll discuss the problem with application state later.
* While scaling resources is eventually often inevitable, it's often better and cheaper to resist the temptation to throw hardware at the problem and postpone scaling events as long as possible by considering runtime efficiency and algorithmic scaling.
* There are various Go features and tools that allow us to identify and fix common problems like memory leaks and lock contention that tend to plague systems at scale.

**Loose Coupling**
* The system property and design strategy of ensuring that a system's components have as little knowledge of other components as possible.
* Distributed monolith: Worst of both worlds. All of the complexities of microservices  + all the tangled dependencies of the typical monolith.

********************************************************************************************************

**Fault Tolerance**
* Synonymous for - self healing, self repair, resilience.
* A system's ability to detect errors and prevent them from cascading to a full blown failure.
* 2 Parts: *Error Detection*, *Recovery*.
* Recovery: System is returned to a state where it can be activated again.
* The most common strategy for providing resilience is redundancy: the duplication of critical components(having multiple service replicas) or functions(retrying service requests).