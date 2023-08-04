**Requirements**
* We can select an item. How to select an item? Each item will have a unique code.
    * There will also be an associated price.
* User first enters money and then selects an item.
* State pattern helps us in understanding the kinds of operations that we can do for a given state. Eg:
    * When we are in idle state, we can only perform one operation:
        * press insert cache button.

**How to implement state design pattern?** 
* Create an interface which has all the methods.
* Each state will either implement the method or throw an error.
* There will be a context struct. 
* Each state implementation will have access to an instance of context struct to change the state.

**Vending Machine**
* Important that all input and output is done through this VendingMachine struct which has all the necessary functions.
* This will do all the initializations so that rather than call state.func we call VendingMachine.func
* All the state structs will be initialized inside it only and need not be initialized explicitly.
* We can now also smoothly transition from one state to another using:
    * state.VendingMachine.setState(state.VendingMachine.noMoney)

**Best Practice**
* Keep on testing each and every function separately.
* Focus on input and output and making it easy for the user.