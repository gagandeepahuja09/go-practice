/*
* when the object can be in many different states and depending on the current request
* the object changes.
* eg. of Vending machines
* When the vending machine is in "itemRequested" state then it will move to "hasMoney"
* state when the action "insertMoney" is done.
* breaking our requirements down to the required ACTIONS and STATES is key for state
* pattern.
*
*
* used when an object can have different responses to the request depending upon the
* current state.
* eg. on purchasing, we will have different responses depending on if it is in hasItem
* or noItem state.
* NO CONDITIONAL STATEMENTS for these. All the logic is handled by state implementation.

* UML Diagram:
* State interface is embedded inside Context struct
* Concrete state 1, 2, 3, 4 implement the state interface

* Let's assume that there are 4 states and 4 actions
* Each of the state would be a struct and each of the action would be a method
* part of the interface that all the structs need to implement.
* Concrete State 1 ===> noItem
* Concrete State 2 ===> hasItem
* Concrete State 3 ===> itemRequested
* Concrete State 4 ===> hasMoney

* Actions or methods to be implemented by each of the structs:
* requestItem
* addItem
* dispenseItem
* insertMoney
 */