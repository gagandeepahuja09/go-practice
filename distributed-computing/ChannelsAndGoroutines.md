Concurrency And Parallelism

* Concurrency: Dealing with a lot of things at once. However, we will be doing only one thing at a time. eg. One task is waiting and the program decides to run another task in the mean time.

* Parallelism: Doing a lot of things at once.

* Serial With Goroutines: We added go keyword to all the split tasks. We saw that the output for methods like continueWritingMail1, continueWritingMail2, continueListeningToAudioBook is missing.
* The reason being that we are using goroutines.
* Since goroutines are not waited upon, the code in the main function keeps on executing and once the control reaches the end of the main function, the program ends.
* What we want to do is to wait in the main function until all goroutines have finished executing.
* Two ways to do this - using channels or Waitgroup.

Waitgroups
* Waitgroup.Add(int) ==> To keep a count of how many goroutines we will be running as part of our logic.
* Waitgroup.Done()   ==> To signal that a goroutine is done with its task.
* Waitgroup.Wait()   ==> To wait until all goroutines are done.
* Pass Waitgroup instance to the goroutines so that they can call the done method.