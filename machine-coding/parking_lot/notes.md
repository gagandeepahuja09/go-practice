* ChatGPT had built a design better than me.
* I had kept slot related information as a property of car itself while we should decouple that and keep that information in the parking lot entity itself as Car will keep on coming and going.
* []*car datastructure was used. This helped in keeping the slots sequential and easily fitting a new car if needed. Nil car pointer at a slot will indicate that it is not occupied.
* *Strategy Pattern*: Parking strategy: NearestAvailableStrategy, RandomStrategy
* Also implemented Observer pattern.