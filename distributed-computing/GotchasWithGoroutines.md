Single Goroutine halting the entire execution.
* Go allows us to recover from a panic using the recover method. We need to ensure that we handle the recover via defer along with handling wg.Done() in defer.

Goroutines aren't predictable
* From parallelism example, we can infer that there were atleast 2 Ps.
* Consider cases where this might not be true
    * GOMAXPROCS = 1
    * Low hardware capabilities.
* We shouldn't rely on them to be executed in any particular order.