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