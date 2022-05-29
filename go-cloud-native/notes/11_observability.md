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
* Traditionally monitoring focuses on asking questions in the hope of identifying or predicting some *expected or previously observed failures*.
    * Assumption is that the system is expected to fail in a specific, predictable way.
    * When a new failure mode is discovered(usually the hard way) - its symptoms are added to the monitoring suite and the process begins again.
    * Problems:
        * Asking new questions of a system often means writing and shipping new code. Not flexible, scalable + super annoying.
        * At a certain level of complexity, the no. of "unknown unknowns" in a system start to overwhelm the no. of "known unknowns".
* Monitoring is something that you do to find out if a system isn't working. Observability is a property a system has that lets you ask why it isn't working.

**The 3 Pillars Of Observability**

* *Tracing(or Distributed Tracing)*
    * Follows a request as it propagates through a distributed system, allowing the entire end-to-end request flow to be reconstructed as a DAG called a *trace*.
    * Analysis of these traces can provide insights into how a system's components interact, making it possible to pinpoint failures and performance issues.

* *Metrics*:
    * Collection of numerical data points representing the state of various aspects of a system at specific points in time.
    * Can be useful to highlight trends, identify anomalies, and predict future behaviour.

* *Logging*:
    * Appending records of noteworthy events to an immutable record - the log - for later review and analysis.
    * Can take various forms, from a continuously appended file on disk to a full-text search engine like ES.
    * Valuable, context-rich insights application-specific events emitted by processes.
    * Not having structured logs can largely reduce their utility.

**The (So-Called) 3 Pillars**
* Just having logging, metrics, and tracing won't necessarily make a system more observable. 