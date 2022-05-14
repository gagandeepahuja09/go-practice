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