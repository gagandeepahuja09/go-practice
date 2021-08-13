// allows you to change the behaviour of an object at run-time without any change
// in the class of that object.

// eg. cache could have various different eviction algorithms.
// we should be able to change algorithm at run-time.
// also cache class should not need to change when a new algorithm is added.

// each of the eviction algorithm will have it's own separate class
// and will implement the evictionAlgo interface.

// when to use
// 1. when an object needs to support different behaviour and we need to change
// the behaviour at run time.
// 2. when you want to avoid a lot of conditionals of choosing the runtime behaviour
// 3. when we have diff algos that are similar and only differ in the way they
// they execute some behaviour.

// setStrategy(setEvictionAlgo) allows to change 