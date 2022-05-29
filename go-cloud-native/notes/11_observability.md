* Data is not information, Information is not knowledge, Knowledge is not understanding, Understanding is not Wisdom.
* The functions of the network and hardware layer are being increasingly abstracted and replaced with API calls and events.
* While we sacrifice a fair share of control over the platforms our software runs on, we win big in overall manageability and reliability, allowing us to focus our limited time and attention on our software.

* This also means that most of our failures now originate from within our own services and the interactions between them.
* No amount of fancy frameworks or protocols can solve the problem of bad software. A kludgy application in k8 is still kludgy.

**What is Observability**
* System property of how well a system's internal states can be inferred from knowledge of its external outputs.
* A system can be considered observable when it's possible to quickly and consistently ask questions about it with minimal prior knowledge and without having to reinstrument or build new code.
* Observability is more than tooling.
* We have to embrace the fact that we can't fully understand a system's state at a given snapshot in time. Understanding all possible failure states in a complex system is pretty much impossible.

**Why do we need observability?**
* Problems that observability tries to address:
    * How do we monitor distributed systems given the ephemerality of modern applications and the environment in which they reside.
    * How can we pinpoint a defect in a single component within the complex web of higly distributed system.

**How is observability different from traditional monitoring**