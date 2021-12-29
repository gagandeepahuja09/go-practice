* The runtime would request the OS to start an ample no. of machines(M), GOMAXPROCS no. of processors(P) to execute goroutines(G).
* Imp: M is the actual unit of execution and G is the logical unit of execution.
* Let's take an example
* We have M1...Mn ready to ready to run.
* 2 Ps, P1 and P2 with runq1, runq2.
* 20 Gs.
* The global scheduler would ideally distribute the goroutines b/w the 2 processors equally.
* Assume P1 is assigned G1...G10 and P2 is assigned G11...G20 and put them in their runqueue.
* P1 runs G1 on M1 and P2 runs G11 on M2.
* [Diag] [https://docs.google.com/document/d/113yv_I2402dnmsKg_8qd82Bn3DnQUBANDpNYYolg7xY/edit]
* A process's internal scheduler is also responsible for switching out current goroutine with the next one it wants to execute. Due to following reasons:
    * Time slice for current execution is over: Put back in runq.
    * Done with the execution: Not put back.
    * Waiting on system call: 
        * The G will be blocked due to this.
        * The P is not req. to wait on the sys call and can leave the waiting G & M combo, which will be picked up by the global scheduler after the sys call. 
        * In the meantime P can pick another M from available machines and another G from runq.

Work-stealing strategy
* If P1's runq is now empty and P2 still has goroutines.
* P1 starts checking with other processors and if another process has G in its runq, it will steal half of them and start executing them.
* What is a processor realizes that it can't steal any more tasks?
    It will wait for a while expecting new goroutines and if none are created, the processor is killed.