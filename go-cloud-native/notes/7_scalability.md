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
    * Offloading the data to an external dedicated cache(at the expense of network I/O.
    * Simply increasing the amount of available memory.

*Disk I/O*
* Scaling strategies:
    * Caching data in RAM(at the expense of memory).
    * Using an external dedicated cache.(expense of network I/O).

*Network I/O*
* Translates directly into how much data the network can transmit per unit of time.
* Scaling strategies for network I/O are often limited.

**************************************************************************************

**State And Statelessness**

* We shouldn't store any state in the application server.
* How can a service then implement in-memory caching. Via pub-sub listen to any changes done and write to all servers.
* State: Set of an application's variables which if changed affect the behavior of the application. 

**Application State Vs Resource State**
* *Application State*: When an application needs to remember an event locally.
* *Resource State*: It is the same for every client and has nothing to do with the action of clients.
* Saying that an application is stateless doesn't mean that it doesn't, just that it's designed in such a way that it's free of any local persistent data.

* Multiple instances of a stateful service will quickly find their individual states diverging due to different inputs being received by each replica.
* Server affinity provides a workaround for this, but it can pose considerable data risk, since the failure of a single server is likely to result in loss of data.

**Advantages Of Statelessness**

* Scalability
    * Each server can handle any request, allowing applications to grow, shrink, or be restarted without loosing data required to handle any in-flight session or requests.
    * Very important for autoscaling, because the pods can(and will) be created and destroyed unexpectedly.
* Durability
    * Data that lives in only one place can get lost when that replica goes away for any reason.
* Simplicity
    * Don't require maintaining service-side state synchronization, consistency and recovery logic makes stateless APIs less complex, hence easier to debug, maintain and build.
* Cacheability  
    * If a service knows that the result of certain request will be same, irrespective of which replica makes the request, cacheability becomes easier.

**Scaling Postponed: Efficiency**

* Efficiency: Ability of a system to handle changes in demand without having to add(or greatly over-provision) dedicated resources.
* If a language has a relatively high per-process concurrency overhead - often the case with dynamically typed languages - it will consume all available memory or compute resources much more quickly than a light-weighed language, and require resources and more scaling events to support the same demand.
* Why are dynamically typed languages slow?
    * Because they must make all their checks at runtime and type is not known to us at the start.

**Efficient Caching Using An LRU Cache**

* Requirements for Cache:
    * Supports *concurrent* read, write and delete operations.
    * Scales well as the number of *cores and goroutines increase*.
    * Won't grow wihout limit to consume all available memory.
* hashicorp/golang-lru ==> well documented, includes sync.RWMutex for concurrency safety.

* Hashicorp's library contains 2 construction functions: 

// New creates an LRU Cache with the given capacity
func New(size int) (*Cache, error)

// NewWithEvict creates an LRU Cache with the given capacity, and also accepts an
// "eviction callback" function that's called when an eviction occurs.
// Can be used for saving somewhere else, like in disk.
func NewWithEvict(size int, ) (*Cache, error) 

* Common methods: 

func (c *Cache) Add(key, value interface{}) (evicted bool)
func (c *Cache) Contains(key interface{}) bool
func (c *Cache) Get(key interface{}) (value interface{}, ok bool)
func (c *Cache) Len() int
func (c *Cache) Remove(key interfaceP{}) (present bool)

* Limitation of LRU: At very high levels of concurrency - on the order of several millions operations per second, it will start to experience some contention.

* What is contention? 
    * A conflict over access to a shared resource.

**Efficient Synchronization**

* Don't communicate by sharing memory; share memory by communicating.
* Channels + Goroutines can allow locks to be dispensed with altogether.
* Channel shine when:
    * Working with many discrete values.
    * Passing ownership of data.
    * Distributing units of work.
    * Communicating async results.
* Mutexes:
    * Synchronizing access to caches or other large stateful structures.

**Share Memory By Communicating**

* Threading is easy; locking is hard.
* Example to show how channel can make concurrency easier compared to locks.
* A program that polls a list of URLs.

**Reduce Blocking With Buffered Channels**

* In case of unbuffered channels, every send blocks until there is a corresponding receive.
* Sends from a channel only block when a buffer is full and receives only block when a buffer is empty.
* Buffered channels are especially useful for handling "bursty" loads.
* In FileTransactionLogger implementation, if events was a standard channel, each send would block until the receiving goroutine method completed a receive operation.
* That might be fine most of the time, but if several writes came in faster than the goroutine could actually process them, the upstream client would be blocked.
* With size of 16, the 17th closely clustered write would block.

* Using buffered channels creates a risk of data loss should the program terminate before any consuming goroutines are able to clear the buffer.

**Minimizing Locks With Sharding**

* Channels can't solve every concurrency problem.
* Example: A large, central data structure such as cache, that can't be easily decomposed into discrete units of work.
* When shared data structures have to be concurrently accessed, it's standard to use locking mechanism such as mutex. eg. a struct that contains a map and an embedded sync.RWMutex.

var cache = struct {
    sync.RWMutex
    data map[string]string
}{ data: make(map[string]string)}

func ThreadSafeWrite(key, value string) {
    cache.Lock()
    cache.data[key] = value
    cache.Unlock()
}

* Lock Contention: As the no. of concurrent processes acting on the data increases, the avg amount of time that the process spends waiting for the locks to be released also increases.
* This could be solved by scaling the no. of instances of the cache service. But this would increase the complexity and latency, as distributed locks needs to be established and writes need to establish consistency.
* Vertical sharding: Only a portion of the overall structure needs to be locked at a time, decreasing the overall lock contention.

***************************************************************************************************

**Memory leaks can... fatal error: runtime: out of memory**

* Memory leaks: Class of bug in which memory is not released even after it's no longer needed.
* More subtle in languages like C++ where memory is manually managed.
* Garbage collected languages like Go aren't immune to memory leaks:
    * Data structures can still grow unbounded.
    * Unresolved goroutines can still accumulate.
    * There could be unstopped time.Ticker values.

**Leaking Goroutines**
* Goroutines are the single largest source of memory leaks in Go(no data to prove it).
* Whenever a goroutine is executed, it's initially allocated a small memory stack - 2048 bytes - that can be dynamically adjusted up or down as it runs to suit the needs of the process.
* The maximum stack size is reflective of the amount of available physical memory.
* When a goroutine returns, its stack is either deallocated or set aside for recycling.
    * Whether by design or by accident, not every goroutine actually returns. Eg. a goroutine is reading from a channel to which we don't write anything.
    * If such a function is called regularly the total amount of memory consumed will slowly increase over time until it's completely exhausted.
    * It's hard to know whether such a process was created intentionally.
* Dave Cheney advice: 
    * You shouldn't start a goroutine without knowing how it will stop.
    * If you don't know how and when a goroutine will exit, that's a potential memory leak.

**Forever ticking tickers**
* time.Ticker => fires repeatedly at some specified interval.
    * Why can't we use time.Sleep instead? We have a better control on closing this via listening to done channel.
* time.Timer  => fires at some point in the future.
* Problem: running time.Ticker values contain an active goroutine that can't be cleaned up.
* By calling ticker.Stop(), we shut down the underlying ticker allowing it to be recovered by the garbage collector.

***************************************************************************************************

**Service Architectures**

**Monolith Service**
* Problems: 
    * Making even a small change requires a new version of the entire monolith to be built, tested and deployed.
    * Cascading failures.
    * Despite the best of intentions and efforts, monolith code tends to decrease in modularity over time, making it harder to make changes in one part of the service without affecting another in unexpected ways.
    * Scaling the application means creating replicas of the entire application, not just the parts that need it. 