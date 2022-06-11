**Manageability**
* Ease with which a system's behavior can be modified to keep it secure, running smoothly and compliant with changing requirements.
* It goes far beyond configuration files.

**Manageability is not Maintainability**
* Both have some mission overlap. They both are concerned with the ease with which a system can be modified.

* *Manageability*:
    * Make changes easily, without having to change the code.
    * It's how easy it is to change a system from the outside.

* *Maitainability*:
    * How easy it is to make changes or add capabilities, correct faults or defects, or improve performance usually by changing code.
    * It's how easy it is to change a system from the inside.

**What is Manageability and Why Should I Care?**
* We should not focus on manageability only in terms of a single service.
* For a system to be manageable, the entire system has to be considered.
    * Can its components be modified independently of one another?
    * Can they be easily replaced if necessary?

* 4 Key functions covered by manageability: 

* *Configuration and Control*:
    * Each of a system's component should be easily configurable.
    * Some systems need regular or real-time control, so having the right "knobs and levers" is absolutely fundamental.

* *Monitoring, logging, and alerting*

* *Deployments and updates*
    * The ability of a system to deploy, update, roll back, and scale system components.
    * This comes into effect throughout a system's lifetime.
    * Lack of external runtimes and singular executable artifacts makes this an area in which Go excels.

* *Service discovery and inventory*
    * Components should be able to quickly and accurately detect one another.

* Managing complex systems is generally difficult and time consuming.
* The costs of managing a system can far exceed the costs of the underlying hardware and software.
* Apart from management costs, manageability will also provide complexity reduction making it easier and faster to undo when it inevitably creeps in.
* Hence it directly impacts reliability, availability and security.