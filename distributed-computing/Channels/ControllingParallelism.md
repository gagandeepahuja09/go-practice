* Spawned goroutine will start executing asap and in a simultaneous fashion. There is an inherent risk involved with that.
* The goroutines might be working on a common source that has a limit on the no. of simultaneous tasks that it can handle.
* This might cause the common source to slow down / fail / produce unexpected output.
* Example: A cashier who has to process orders but has a limit of only 10 orders per day.
* We'll see an example where multiple goroutines are trying to access the variable ordersProcessed and the condition ordersProcessed < 10. This could lead to a race condition, multiple variables trying to access the variable at the same time and that in turn affecting the resource.
* Possible ways to handle the limit:
    * Inc. limit of processing orders ==> can only be done to an extent.
    * Inc. no. of cashiers.


Distributed Work without Channels
* This will still lead to race conditions.
* The variable ordersProcessed changed from 2 to 1 in the logs. This shouldn't have happened.
* Using channels will help us solve the race condition. Using mutexes, semaphores, locks is also an option.
* Channels provide us with much more versatility in terms of usage and provide a larger layer of abstraction(as we need not use lock and unlock methods).
* go run -race main.go


Distributed Work With Channels
Step
* Create channel to accept the data to be processed.
* Launch the goroutines that are waiting on the channel for data.
* We can pass the data to the channel through main goroutine or any other goroutine.
* The goroutines listening to the data can accept the data and process them.
* The advantage of using channels is that multiple goroutines can wait on the same channel and execute tasks concurrently.