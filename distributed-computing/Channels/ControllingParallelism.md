* Spawned goroutine will start executing asap and in a simultaneous fashion. There is an inherent risk involved with that.
* The goroutines might be working on a common source that has a limit on the no. of simultaneous tasks that it can handle.
* This might cause the common source to slow down / fail / produce unexpected output.
* Example: A cashier who has to process orders but has a limit of only 10 orders per day.
* We'll see an example where multiple goroutines are trying to access the variable ordersProcessed and the condition ordersProcessed < 10. This could lead to a race condition, multiple variables trying to access the variable at the same time and that in turn affecting the resource.