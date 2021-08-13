package main

/*
* go doesn't have keywords like extends and implements.
* go prefers composition over inheritance.
* It allows embedding of struct into another struct.
* We can pass the child type in methods which expect the base the type
* There is not type inheritance in go, which means that you cannot pass child type
to a function that expects the base type.
*/
