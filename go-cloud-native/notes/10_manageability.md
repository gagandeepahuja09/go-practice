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

**Decoding arbitrary JSON**

**Field formatting with struct field tags**
* Under the covers, marshalling works by using reflection 
* Current problems 
    * What happens if we marshal struct field tags? We get a json with all the zero values.
    {"Host":"", "Port": 0, "Tags": null}
    * The marshalled response has uppercase characters due to them being exported fields.
* Solution to both problems: struct field tags: most commonly used by encoding packages.
* These tags can be accessed using run-time reflection via the reflect package.

* *Customizing JSON Key*
* *Omitting Empty Values*: the fields will be skipped if they contain a zero-value.
* *Ignoring a field*: Field using the - (dash) option will be completely ignored during encoding and decoding.

type Tagged struct {
    CustomKey string `json:"custom_key"`
    OmitEmpty string `json:",omitempty"`
    IgnoredName string `json:"-"`
    TwoThings string `json:"two_things,omitempty"`
}

**Working with YAML(YAML Ain't Markup Language)**
* It's very popular with projects like Kubernetes that depend on complex, heirarchical configurations.
* Configurations that use it can start to suffer from readability issues as they scale up.
* Unlike JSON which was created as a data-interchange format, YAML is largely a configuration format.
* Advanatages of YAML over JSON:
    * It can self-reference.
    * It allows embedded block literals.
    * Supports comments and complex data-types.

**Encoding YAML**
* Using go yaml. Method signatures exactly same as encoding/json.
* func Marshal(v interface{}) ([]byte, error)
* Just like the version provided by encoding/json, go-yaml's marshal function traverses the value v recursively. Any composite types that it finds - array, slices, maps, structs - will be encoded appropriately and will be present in the output as nested YAML elements.

**Decoding YAML**
* func Unmarshal(data []byte, v interface) error
* If v is nil or not a pointer, yaml.Unmarshal returns an error.

**Watching for configuration file changes**
* Situations where changes need to be made to the configuration of a running program.
* If it doesn't explicitly watch for and reload changes, then it will have to be restarted to reread its configuration.
    * Which can be inconvenient at best and introduce downtime at worst.
* Restarting the program when the configuration changes is also fairly common.

**Making our configuration reloadable**
* We will want to have a single global instance of our configuration struct.
* In a slightly larger project, we might even put this is in a separate config package.
    var config Config

* Antipattern: Configuration being passed to every method. Also, since configuration now lives in N places instead of one, it makes configuration reloading more complicated.

**loadConfiguration Method**
* We created the global config variable.
* loadConfiguration(filepath string) (Config, error)
    * File is read using ioutil.ReadFile which will return the byte slice and an error.
    * Check for errors.
    * Unmarshal the byte slice to an empty variable of type config.
    * Check for error and returned the result config after unmarshal.

**startListening Method**
* startListening accepts two channels: updates which emits the name of the file when that file changes and an errors channel.
* It watches both channels in a select inside an infinite loop.
* Whenever a file updates, the updates channel sends its name, which is then passed to load configuration.
* The configu struct returned by loadConfiguration replaces the current value of global config.

**init Method**
* It will retrieve the channels from a watchConfig method and pass them to startListening.

**Polling for configuration changes**
* Using a time.Ticker to recalculate a hash of the config file every few seconds and reload if the hash changes.

* Go makes a number of common hash algorithms available in its crypto package, each of which lives in its own subpackage of crypto and satisfies both the crypto.Hash and io.Writer interface.

* Eg. crypto/sha256. 
    * sha256.New() to get a new sha256.Hash value.
    * Into which, we write the data we want to calculate the hash of, just a we would any io.Writer.

* Generating hash of a configuration has 3 distinct parts: 
    1. Getting a []byte source in the form of an io.Reader.
    2. We copy those bytes from the io.Reader to our sha256.
    3. Retrieve the hash sum from hash.

**watchConfig method**
* It will return the read only channels (changes, errors, any error in this method).
* Using time.Ticker, every second it will calculateFileHash and send the result/err to the desired channel. 
* All this will be done in a separate goroutine.

**Polling approach pros and cons**
*Pros*:
    1. Straightforward.
    2. Since hashing only cares about the configuration's contents, it can even be generalized to detect changes in places like remote key/value stores which technically aren't files.

*Cons*:
    1. Can be computationally wasteful, especially for very large or many files.
    2. Brief delay between the time the file is changed and the detection of the change.