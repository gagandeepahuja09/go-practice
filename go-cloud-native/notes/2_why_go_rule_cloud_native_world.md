Features for a Cloud Native World

* Other languages before Golang: overly complex.

Issues
1. Low Program Comprehenisibility: Cleverness over clarity was becoming common.
2. Slow build times: Language construction and years of feature creep.
3. Inefficiency: Dynamic languages trading efficiency and type safety for expressiveness.
4. High cost of updates: Incompatibility b/w even minor versions of a language.

Composition and Structural Typing

Comprehensibility

CSP-Style Concurrency

Fast Builds
* Slower builds => Loose out on developer productivity.
* Go => Not having complex relationships b/w files, simplifies the dependency analysis.
* Reduces the overhead of C-style include files and libraries.
* Possible cons => some promising Go features getting skipped due to affecting the build times.

Linguistic Stability
* Programs written in Go 1 will containue to compile and run correctly, unchanged, for the lifetime of Go 1 specification.

Memory Safety
* Go neither needs nor allows the kind of manual memory management and manipulation that lower-level languages like C, C++ allow and require.
* This decreases complexity and increases memory safety.
* Pointers are strictly typed and always initialized to some value, even if that value is nil.
* Pointer arithmetic is strictly disallowed.
* Built-in reference types like maps and channel, which are internally represented as pointer to mutable structures are initialize by the make function.
* By reducing manual memory management, it has become possible to remove entire class of memory errors and security loops.
    * No memory leaks, no buffer overruns, no address space layout randomization.
* No need to track and free up memory for every allocated byte.  
* Tradeoffs of all this: Can't compete with C++, Rust in pure raw execution speed.

Performance

Static Linking
* Static linking: Copying all library modules used in the program into the final executable image.
* Dynamically linking: Loading the external shared libraries into the program and then bind those shared libraries into the program.
* Static linking produces slightly larger binary executable files. (of the order of few MBs).
* Pros:
    * Resulting binary has no external language runtime to install.(happens in case of Java)
    * No external libary dependencies to upgrade or conflict.(happens in case of Python).
* It can be easily distributed to users or deployed to a host without fear of suffering dependency or environment conflicts. This ability is particularly useful when working with containers.
* Because of no dependencies, go binaries can be built into scratch images that don't have parent images.
* This helps with minimal deployment latency and data transfer overhead.
* These are very useful traits in an orchestration system like Kubernetes that may need to pull images with some regularity.

Static Typing
