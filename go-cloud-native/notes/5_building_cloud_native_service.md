* Building a distributed key value store.

Requirements 
* Store key-value pairs.
* Service endpoints.
* Persistence.
* Idempotence

* Idempotence: Calling a function once has the same effect as calling it multiple times.
* Nullipotence: Trigger no state change. Eg. x = 1, HTTP PUT is idempotent but not nullipotent. HTTP GET, x=x are nullipotent as well.

Why do we need idempotence?
* Safer: Retries would give unexepected results in not idempotent.
* Idempotent operations are often simpler: Eg. a PUT method requiring only setting value vs. a POST method that checks whether a value is set and depending on that throwing an error.
* Idempotent operations are declarative 
    * They tell a service what needs to be done(declarative), instead of telling how to do it(imperative).
    * This frees users from dealing with low-level constructs and minimizes potential side-effects.
* Cloud-native services must hold this property.

************************************************************************************

Generation 0: The Core Functionality
* While returning the error, we don't use errors.New. Instead we return the prebuilt ErrNoSuchKey error value.
* This is an example of sentinel error, which determines exactly what type of error it's receiving and to respond accordingly.

    if err.Is(ErrNoSuchKey) {
        http.Error(w, err.Error(), http.StatusNotFound)
        return
    }

***********************************************************************************

Generation 1: The Monolith
* REST => Perfectly sufficient for our needs.
* Go standard libraries are designed to be extensible, so there are a no. of web frameworks that extend them. 
* Handler => any type that implements the Handler interface.
* mux/multiplexer => direct incoming signals to one of the possible outputs.
* When a request is received by a service that's been started by ListenAndServe, it's the job of a mux to compare the request URL to the registered patterns and call the handler function with the one that matches closely.
* If mux is specified as nil, the DefaultServeMux is used.

gorilla/mux
* For many of the use-cases, DefaultServeMux should be fine.
* One of the key features of gorilla/mux is the ability to create paths with variable segments which can optionally contain a regular expression pattern.
* Format => {name} or {name:pattern}
    r := mux.NewRouter()
    r.HandleFunc("/products/{key}", ProductHandler)
    r.HandleFunc("/articles/{category}/", ArticlesCategoryHandler)
    r.HandleFunc("/articles/{category}/{id:[0-9]+}, ArticleHandler)
* We can retrieve the variables using mux.Vars function which returns a map[string]string.
* Matchers => We can add additional matching criterias. Eg domains/subdomains, methods, schemes, path prefixes, headers or even custom matching functions that we create. 
    * Each matcher returns a route, hence can be easily chained.
    * r.HandleFunc("/products", ProductHandler).
        Host("www.example.com").
        Methods("GET", "PUT").
        Schemes("http")

Building a RESTful service