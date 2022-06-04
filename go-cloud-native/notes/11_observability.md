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

**OpenTelemetry**
* Unlike most CNCF projects, it isn't a service per se.
* It's an effort to standardize how telemetry data - traces, metric and eventually (logs) are expressed, collected, and transferred.
* There are many vendors and tools in the instrumentation space.
* OpenTelemetry seeks to unify this space by a vendor-neutral specification that standardizes how telemetry data is collected and sent to backend platforms.

**The OpenTelemetry Components**
* OpenTelemetry extends and unifies earlier attempts at creating telemetry standards, by including abstractions and extension points in the SDK where you can insert your own implementation.
* This makes it possible to implement custom exporters that can interface with a vendor of your own choice.

* Core Components:
* *Specifications*: These describe the requirements and expectations for all OpenTelemetry APIs, SDKs, and data protocols.

* *API*: Language specific interfaces and implementations based on the specifications that can be used to add OpenTelemetry to an application.

* *SDK*: 
    * The concrete OpenTelemetry implementations that sit between the APIs and the Exporters, provide functionality like (for example) *state tracking* and *batching data* for transmission.
    * It also provides a no. of configuration options for behaviors like *request filtering* and *transaction sampling*.

* *Exporters*:
    * In-process SDK plug-ins that are capable of sending data to a specific destination which may be local(log file or stdout) or remote(Jaeger or some commercial solution).
    * Exporters decouple the intrumentation from the backend, making it possible to change destinations without having to reinstrument the code.

* *Collector*:
    * Optional but very useful vendor-agnostic service that can receive and process telemetry data before forwarding it to one or more destinations.
    * Can run as a sidecar process alongside your application or a standalone proxy elsewhere.
    * Can be particularly useful in the kind of tightly controlled environments that are common in the enterprise.

* OpenTelemetry is only concerned with the collection, processing, and sending of telemetry data, and relies on you to provide a telemetry backend to receive and store the data.

*************************************************************************************

**Tracing**

* It's often the challenge to find the source of the problem before we can actually fix this.
* Tracing helps to solve this problem by tracking requests as they propagate through the system - even across process, networks, and security boundaries.
* Tracing can help:
    * Pinpoint component failure.
    * Identify performance bottlenecks.
    * Analyze service dependencies.
* Tracing is generally discussed in the context of distributed systems, but a complex monolith application can also benefit from tracing, especially if it contends for resources like network, disk or mutexes.

**Tracing Concepts**

* *Spans*: 
    * Describes a unit of work performed by a request.
    * Eg. a hop across the network or a fork in the execution flow.
    * Each span has an associated name, a start and a duration.
    * They can be and typically are nested and ordered to model a causal relationship.

* *Traces*:
    * Represents all of the events - individually represented as spans - that make up a request as it flows through a system.
    * DAG of spans or stack trace where each span represents the work done by one component.

* When a request begins in the first(edge) service, it creates the first span - the root span.
* Root span is automatically assigned a globally unique trace ID, which is passed with each subsequent hop in the request lifecycle.
* The next hop can either choose to insert or otherwise enrich the metadata associated with each request.

**Tracing with OpenTelemetry**
* 2 phases: configuration & instrumentation.
* This is common for both tracing & metrics.
* The configuration phase is executed exactly once in a program, usually in the main function.
* *Configuration steps*: 
    1. *Retrieving and configuring the appropriate exporters* for your target backends.
        * Tracing exporters implement the SpanExporter interface(OpenTelemetry).
        * Also includes several stock exporters by OpenTelemetry.
    2. *Passing the exporters and any other appropriate configuration options to the SDK to create the *tracer provider* *. 
        * Tracer provider will serve as the main entry point for the open telemetry tracing API for the lifetime of our program.
    3. *Setting the global tracer provider*.
        * This make it discoverable via the otel.GetTracerProvider function.
        * It allows the libraries and other dependencies that also use the OpenTelemetry API to more easily discover the SDK and emit telemetry data.

* *Instrumenting code steps*:
    1. *Obtaining a Tracer*
        * Tracer interface: It has the central role of keeping track of trace and span information from the (usually global) tracer provider.
    2. *Starting and ending spans*
    3. *Setting span metadata*
        * Includes timestamped messages called events or key/value pairs called attributes.


**Creating the tracing exporters**
* Tracing exporters implement the SpanExporter interface.
* Which in otel 0.17.0 lives in go.opentelemetry.io/otel/sdk/export/trace aliased as export.
* OpenTelemetry exporters are in-process plug-ins that know how to convert metric or trace data and send it to a particular destination.
* The destination may be local(stdout or a log file) or remote(Jaeger or a commercial solution).

* *The Console Exporter*
    * Write telemetry data as JSON to stdout.
    * Console exporter can be used to export metric telemetry.
    * Package: go.opentelemetry.io/otel/ exporters/stdout package.
    * Like most exporters' creation functions, stdout.NewExporter is a variadic function that can accept zero or more configuration options.
        stdExporter, err := stdout.NewExporter(
            stdout.WithPrettyPrint(),
        ) 

* *The Jaeger Exporter*
    * It knows how to encode tracing telemetry data to the Jaeger distributed tracing system.
    * https://www.jaegertracing.io/
    * https://eng.uber.com/distributed-tracing/
    * Jaeger key features:
        * Multiple storage backends: Cassandra, ElasticSearch, Kafka, memory.
        * Opentracing inspired data model.
    * Both stdout.NewExporter and jaeger.NewExporter return an export.SpanExporter.

    jaegerEndpoint := "http://localhost:14268/api/traces"
    serviceName := "fibonacci"

    jaegerExporter, err := jaeger.NewRawExporter(
        jaeger.WithCollectorEndpoint(jaegerEndpoint),
        jaeger.WithProcess(jaeger.Process{
            ServiceName: serviceName,
        }),
    )

**Creating a tracer provider**
