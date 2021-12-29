* Go runtime uses a scheduler strategy known as M:N scheduler which will schedule M no. of goroutines on N no. of OS threads.
* This provides faster context swith b/w goroutines. 
* Also enables us to use the multiple cores of CPU for parallel computing.

* 3 Main entities from go runtime perspective
* G => Goroutine
* M => OS Thread or Machine
* P => Context or Processor


Goroutine
* It is the logical unit of execution that contains the actual instructions for our program/functions to run.
* It also contains other important information like: the stack memory, which machine(M) is it running on, and which go function called it.
* When we start our progam, a goroutine called main goroutine is first launched.
* It takes care of setting up the runtime space before starting our program.
* A typical runtime setup includes things such as: max stack size, enabling garbage collector, etc.

// Denoted as G in runtime
type g struct {
    stack stack // offset known to runtime/cgo
    m *m        // current m;offset know to arm liblink
    goid int64
    waitsince int64     // approximate time when the g became blocked.
    waitreason string   // if status==Gwaiting
    gopc uintptr        // pc of go statement that created this goroutine.
    startpc uintptr     // pc of goroutine function
    timer *timer        // cached timer for time.Sleep
}


OS Thread Or Machine
* Initially the OS thread are created & managed by the OS.
* Later on, the scheduler can request for more OS threads or machines to be created / destroyed.
* It is the actual resource upon which goroutine will be executed.
* It also maintains info about the main goroutine(g0 ==> go routine with scheduling stack), the G currently being run on it(curg), thread local storage(tls), attached p for executing the code, stack that created this thread(createstack), whether it is out of work and is actively looking for work(spinning).

type m struct {
    go *g           // main goroutine
    tls [6]uintptr  // thread-local storage(for x86 extern register)
    curg *g         // current running goroutine. this indicates that in a particular m, there can only be one running go routine at a particular instant
    p uintptr       // attached p for running go code. nil if not running any code.
    id int32
    createstack [32]uintptr // stack that created this thread.
    spinning // m is out of work and is actively looking for work.
}


Context or processor
* We have a global scheduler which takes care of bringing up new M, registering G and handling system calls.
* However, it doesn't handle actual execution of goroutines.
* This is done by an entity called Processor, which has its own internal scheduler and queue called run queue(runq).
* runq consists of goroutines that will be executed in the current context.
* It also handles switching b/w various goroutines.

type p struct {
    id int32
    m uintptr   // back-link to associated m.(nil if idle)
    run [256]guintptr
}

GOMAXPROCS: Max. no of Ps running at any point in a program's life.