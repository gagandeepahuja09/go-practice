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

**Configuring Your Application**
* Anything that's likely to be varied b/w environments - dev, stage, prod.
* 12 factor app - III. Store configuration in the environment.

* *Configuration should be cleanly separated from the code*

* *Configuration should be stored in version control*
    * Storing it in version control, separately from the code allows us to quickly roll back a config change.
    * Deployment frameworks like Kubernetes provide config primitive like ConfigMap for this.

* 3 Common ways to configure applications:
    * Environment variables
    * Command-line flags
    * Configuration files

**Configuration Good Practice**
* *Version control your configurations*
    * Makes it possible to:
        * Review them before deployment.
        * Quickly reference them afterwards.
        * Quickly rollback a change if necessary. 

* *Don't roll your own format*
    * Standard formats: JSON, YAML, TOML.
    * It you must roll your own format, be sure that you're comfortable with the idea of maintaining it and forcing any future maintainers to deal with it forever.

* *Make the zero value useful*
    * Don't use nonzero default values unnecessarily.
    * The behavior that results from an undefined configuration should be acceptable, reasonable and unsuprising.

**Configuring with Environment Variables**
* Merits of using:
    * Env vars are *universally supported*.
    * They ensure that configuration don't get accidentally checked into the code.
    * Generally require less code than configuration files.
    * Perfectly adequate for small applications.

* Demerits:
    * We can't easily learn about the existence and behavior of environment variables by looking at an existing config file. Applications that rely on them can be harder to use and debug.

* name := os.Getenv("NAME"). If variable is not present, Getenv will return an empty string. In order to distinguish between empty value and unset value, we can use os.LookEnv which returns both the value and a boolean.

* For more sophisticated options like default values or typed variables, viper: a 3rd party package is fairly popular. 

**Configuring with Command-Line Arguments**
* They are definitely worth considering, atleast for smaller, less complex applications.
* Merits:
    * They are explicit.
    * They have out-of-the-box type support.
    * They details of their existence and usage are generally available via a --help option.

**The standard flag package**
* Example code: flag.go
* go run . -help --> to see the summary of the flags.
* Problems with flag package:
    * Flag syntax is non-standard. Standard: long form like version with two dashes => --version. short form with single dash => -v.
    * It only parses flags. We can map commands to functions.

**The Cobra command-line parser** 
* It's used in a number of high profile projects like Docker, Kubernetes, Istio, Helm, CockroachDB.
* Merits
    * Provides fully POSIX-compliant flags.(short and long versions).
    * Supports nested subcommands.
    * Automatically generates help output and autocomplete for various shells.
* Demerits
    * Quite complex relative to the flags package.

********************************************************************************

**Configuring with Files**
* Merits:
    * They tend to be more explicit and comprehensible by allowing behaviors to be logically grouped and annotated.
    * Understanding how to use a config file is just a matter of looking at its structure or an example of its use.
    * They are particularly useful when managing a large number of options. With command line flags, it can result in pretty long statements.

* Possible demerits:
    * Distributing config files at scale can be a challenge.
        * We can instead use distributed key-value store such as etcd and consul for such use cases.
        * Most orchestraction platforms provide specialized configuration resource like Kubernetes ConfigMap that largely alleviate the distribution problem.

**Our configuration data structure**
* Two general ways in which configuration can be unmarshalled:
    * *Mapping corresponding fields in a specific struct type*. Eg. host: localhost can be unmarshalled into a struct type that has a Host string field.
    * Unmarshalling into one or more possibly nested maps of type map[string]interface{}.
        * This can be convenient when working with arbitrary configurations, but can be very akward to work with.

* If we are aware of what our configuration is going to look like, which is very likely the case, approach 1 is the most appropriate.

type Config struct {
    Host string
    Port uint16
    Tags map[string]string
}

* For a struct to be marshallable or unmarshallable, it must begin with a capital letter to indicate that it's exported by its package.

**Working with JSON**
Demerits:
    * Considered lesser user-friendly than YAML.
    * Unforgiving syntax: can easily be broken by a misplaced or missing comma.
    * Doesn't support comments.

**Encoding JSON**
* It might seem strange but encoding is the first step for understanding how to decode JSON.
* It provides a handy means of generating, testing and debugging our configuration files.
* Go's standard encoding/json package.
    func Marshal(v interface{}) ([]byte, error)
* The JSON.Mashal function traverses the value of v recursively, so any internal structs will be encoded as nested JSON.
* We can also use JSON.MarshalIndent which returns pretty printed JSON. This can be very useful for bootstrapping configuration files.
    json.MarshalIndent(v interface{}, prefix, indent string) ([]byte, error)

**Decoding JSON**
* func Unmarshal(data []byte, v interface{}) error
* json.Unmarshal will only decode the fields that it can find in the target type. The feature can be quite useful if we want to cherry pick a few specific fields out of a big JSON blob. 

**Working with YAML**

**Watching for configuration file changes**