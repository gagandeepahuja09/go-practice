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

********************************************************************************************************

**Fault Prevention**
* Many - if not most - classes of errors and faults can be predicted and prevented during the earliest phases of development.

**Good programming practices**
* Explicit goal of any development methodology is fault prevention - from pair programming, to TDD, to CR practices.

**Language Features**
* Features such as dynamic typing, pointer arithmetic, manual memory management, and thrown exceptions can easily introduce unintended behaviors that are difficult to find and fix, and may even be maliciously exploitable.

**Scalability**
* Ability of a system to continue to provide correct service in the face of significant change in demand.
* We'll discuss horizontal and vertical scaling in detail later.
* We'll discuss the problem with application state later.
* While scaling resources is eventually often inevitable, it's often better and cheaper to resist the temptation to throw hardware at the problem and postpone scaling events as long as possible by considering runtime efficiency and algorithmic scaling.
* There are various Go features and tools that allow us to identify and fix common problems like memory leaks and lock contention that tend to plague systems at scale.

**Loose Coupling**
* The system property and design strategy of ensuring that a system's components have as little knowledge of other components as possible.
* Distributed monolith: Worst of both worlds. All of the complexities of microservices  + all the tangled dependencies of the typical monolith.

********************************************************************************************************

**Fault Tolerance**
* Synonymous for - self healing, self repair, resilience.
* A system's ability to detect errors and prevent them from cascading to a full blown failure.
* 2 Parts: *Error Detection*, *Recovery*.
* Recovery: System is returned to a state where it can be activated again.
* The most common strategy for providing resilience is redundancy: the duplication of critical components(having multiple service replicas) or functions(retrying service requests).

********************************************************************************************************

**Fault Removal**

**Verification and Testing**
* Two common approaches of finding software faults in development:
    * Static Analysis: 
        * Automated, rule-based code analysis performed without actually executing programs.
        * Useful for providing early feedback, enforcing consistent practices, and finding common errors and security holes without depending on human knowledge or effort.
    * Dynamic Analysis/Testing.
* Having software that's designed for testability by minimizing the degrees of freedom - the range of possible states - of its components.

**Manageability**
* A system is said to be manageable if it's possible to sufficiently alter its behavior without having to alter its code.
* Manageability can include the following:
    * Configuration changes.
    * Feature flags.
    * Rotate credentials or TLS certificates.
    * Deploy or upgrade/downgrade system components.
* Manageable systems are designed for adaptability, to accomodate changing functional, environmental or security requirements.
* Unmanageable systems are brittle, frequently requiring adhoc or manual changes.

********************************************************************************************************

**Fault Forecasting**
* Observability, stress testing, failure mode and effects analysis

********************************************************************************************************

**The Continuing Relevance Of the 12 Factor App**
* Developers at Heroku were seeing applications being developed again and again with the same fundamental flaws.
* Methodology was for building applications that: 
    1. Use declarative format for setup automation, to minimize time and cost for new developers joining the project.
    2. Have a clean contract with the underlying OS, offering maximum portability between execution environments.
    3. Are suitable for deployment on modern cloud platforms, obviating the need for servers and systems administration.
    4. Minimize divergence between development and production, enabling continuous development for maximum agility.
    5. Can scale up without significant changes to tooling, architecture, or development practices. 

********************************************************************************************************

**I. Codebase**
* *One codebase tracked in revision control, many deploys.*
* For any given service, there should be exactly 1 codebase that's used to produce any no. of immutable releases for multiple deployments to multiple environments.

* Having multiple services sharing the same code tends to lead to a blurring of the lines, trending in time to something like a monolith, making it harder to make changes in one part of the service without affecting other part or other service.
* Shared code should be refactored into libraries that can be individually versioned and included through a dependency manager.

* Having a single service spread across multiple repos makes it nearly impossible to automatically apply the build and deploy phases of your service's life cycle.

********************************************************************************************************

**II. Dependencies**
* *Explicitly declare and isolate code dependencies*.
* For any given version of the codebase, go build, go test, and go run should be deterministic: they should have the same result however they're run, and the product should respond the same way to the same inputs.
* What if a dependency changes in such a way that it introduces a bug?
    * Most programming languages offer a packaging system for distributing support libraries.
    * Go uses go modules to ensure that imported packages won't change.
* Services should try to avoid using the os/exec package's Command function to shell out to external tools like ImageMagick or curl.

********************************************************************************************************

**III. Configuration**
* *Store configuration in the environment*.
* Anything that's like to vary between environments should always be cleanly separated from the code.
* Configuration files may include but not limited to:
    * URLs or other resource handles to a database or upstream service.
    * Secrets of any kind.
    * Per environment values.
* YAML ==> commonly used configuration.
* Instead of configurations as code or external configurations, it should be stored as environment variables:
    * They are standard and largely OS and language agnostic.
    * They are easy to change between deploys without changing any code.
    * They are very easy to inject into containers.
* Tools in Go for this:
    * os package is the most basic.
        * os.GetEnv("NAME")
    * spf13/viper
        * viper.BindEnv("id")   // will be uppercased automatically
        * viper.SetDefault("id", 13)
        * viper.GetInt("id")    // 13
        * os.SetEnv("ID", "50") // typically done outside of the app.
        * viper.GetInt("id")    // 50
    * Viper provides a no. of additional features: 
        * Default values.
        * Type variables.
        * Reading from command-line flags, variously formatted configuration files, and even remote configuration systems like etcd, Consul.

********************************************************************************************************

**IV. Backing Services**
* *Treat backing services as attached resources*
* Backing service: Any downstream dependency that a service consumes.
* A services should make no distinction b/w backing services of the same type - internal or 3rd party.
* Eg. A MySQL DB run by your own team's sysadmins should be treated no differently than an AWS-managed RDS instance.
* Whether it's running in a data center on another hemisphere or a docker container on the same server.

* A service that's able to swap out resources at will with another one of the same kind - internally managed or otherwise - just by changing a configuration value can be more easily deployed to different environments, tested and maintained. 

********************************************************************************************************

**V. Build, Release And Run**
* *Strictly separate build and run stages*

********************************************************************************************************

**VI. Processes**