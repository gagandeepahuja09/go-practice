Key Learnings:

* How to build complete command-line applications(with as little as single-code file).
* How to ensure that the tools we build can be composed with other tools using standard streams.
* How to interact with a simple 3rd party JSON RESTful API.
* How to utilize the standard in and out packages in Go code.
* How to read from streaming source, one line at a time.
* How to build a WHOIS client to look up domain information.
* How to use and store sensitive or deployment specific information in environment variables.

**************************************************************************

Pipe Design For Command-line tools

* Command-line tools using the standard stream(stdin pipe and stdout pipe) to communicate with the user and with other tools.
* By default, stdin(standard input) is connected to user's keyboard and stdout to the terminal.
* Both can be redirected using redirection metacharacters.
* We can also throw away the output using NUL in windows & /dev/null in Unix machines.
* We can also redirect the output to a file that will cause it to be saved there.
* Alternatively we can pipe the output of one program to the input of another.
* We'll make use of this feature to connect our various tools together.
* eg. echo -n "Hello" | md5 ==> output of Hello is the input of md5 command.
* When we don't use any pipe redirection, we'll able to directly work with programs using the default in and out which will be useful in testinf and debugging.
