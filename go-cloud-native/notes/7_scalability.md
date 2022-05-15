Autoscaling

* Being able to add resources on demand means that we can serve our users under load far beyond what we had ever anticipated.
* If any one server failed, its work can be divide among survivors.
* Having far more resources than necessary is both wasteful and expensive. We need the ability to scale our resources back in when demand ebbed.
* Since this is more of a Go book, focus will be more on the other non-infra part of the scalability equation: efficiency.

**What is Scalability?**
* The ability of a system to continue to provide correct service in the face of significant changes in demand.
* A system can be considered to be scalable if it doesn't need to be redesigned to perform its intended function during steep increase in load.
* This definition doesn't actually say anything about adding physical resources.
* Can be done both by physical resources and via building efficient systems.
* Efficient systems are more scalable because of their ability to gracefully absorb high levels of demand.

**Different forms of scaling**
* Unfortunately, even the most efficient of efficiency strategies has its limit, and eventually you'll find yourself needing to scale your service to provide additional resources.

*Vertical Scaling*
* Scale up.
* Increasing resource allocation by changing its instance size.

*Horizontal Scaling*
* Scale out.
* Duplicating the system or service to limit the burden on any individual server.
* The presence of state can make this strategy difficult or impossible for some systems.

*Functional Partitioning*
* Decomposing complex systems into smaller functional units that can be independently optimized, managed, and scaled.

*Sharding*
* Common in DB.
* Dividing data into partitions, each of which holds a specific subset of the dataset.

**The 4 Common Bottlenecks**
* Common was to solve this is to scale up.
* Other ways: utilizing other resources that the system has in abundance. eg avoiding disk I/O bottlenecking by caching data in RAM. Conversely a memory-hungry service could page data to disk.
* Horizontal scaling doesn't make a system immune: adding more instances can mean more communication overhead, which puts additional strain on the network.

*CPU*
* The number of operations per unit of time that can be performed by a system's central processor.
* Scaling strategies include:
    * Caching the result of expensive deterministic operations(at the expense of memory).
    * Increasing the size/no. of processors(at the expense of network I/O if scaling out).

*Memory*
* The amount of data that can be stored in main memory.
* Scaling strategies:
    * Offloading the data from memory to disk(at the expense of disk I/O).
    * Offloading the data to an external dedicated cache(at the expense of network I/O).
    * Simply increasing the amount of available memory.

*Disk I/O*
* Scaling strategies:
    * Caching data in RAM(at the expense of memory).
    * Using an external dedicated cache.(expense of network I/O).

*Network I/O*
* Translates directly into how much data the network can transmit per unit of time.
* Scaling strategies for network I/O are often limited.