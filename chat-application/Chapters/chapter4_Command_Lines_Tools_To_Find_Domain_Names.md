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


Five Simple Programs
* Sprinkle: Add some web-friendly sprinkle words to increase the chances of finding the available domain names.
* Domainify: Removing unacceptable words for a domain. It will replace spaces with hyphens and add an appropriate top-level domain(.com, .net, etc).
* Coolify: Changing a boring only normal word to web 2.0 by fiddling around with vowels.
* Synonyms: Use 3rd party API to find synonyms.
* Available: Will check to see whether a domain is available or not using appropriate WHOIS server.


Sprinkle
* Define an array of transformations.
* Use bufio package to scan the input from stdin and fmt.Println in order to write the output to stdout.
* Use math/rand to random select a transformation to apply.
* Ideally should be created in GOPATH/src.
computers can't actually generate random no.s But changing the seed gives the illusion that it can. As the seed would be different everytime the program is run.
* Scanner allows to read blocks of bytes separated by defined delimiters, such as carriage return / linefeed characters.
* We can use our own split function for the scanner or use one of the options built-in in the standard lib. eg. bufio.ScanWords.
* Scan tells the scanner to read the next block of bytes. It returns a bool value indicating whether it found a value.
* The bytes that are selected in Bytes method and the Text method gives us the same info converted from byte slice to string.
* math/rand: insecure random numbers. crypto/rand: secure random no.s.
* go build -o sprinkle ==> will build the executable file.
* Can be run using ./sprinkle ==> then keep on typing chat to see different outputs.
* Also we can pipe the result of echo command to our application.
echo "chat" | ./sprinkle


Domainify
* The output from sprinkle may contain spaces and other characters not allowed in domains.
* Domainify will convert a line of text segment into acceptable domain and add an appropriate TLD(top level domain) to the end.
* rune datatype is a type alias of int32, just used to differentiate for character values. 
* To combine the above two applications, go to the parent directory and run the following: ./sprinkle/sprinkle | ./domainify/domainify


Coolify
* When domain names for common words are already taken and it's common to play around with words and add or remove vowels.
* String manipulation using go specific methods like append which use variadic functions.
* Go's method if reprsenting substrings using [start:end]
* Only end mentioned: [:end] ==> prefix
* Only start mentioned: [start:] ==> suffix
* There are 2 cases to handle here: duplicate vowel or remove vowel. Will use random no. to decide on that.
* We will use append method to do the string manipulation. append has following parameters: first is the original string and the second are n no.s in order which we want to append.
* For remove vowel, it means the we have 0...i-1 followed by i+1...end ==> this means first parameter would be v[:i-1] and the second one would be all no. from i+1...end. We can rather express it in variadic form as v[i+1:]...
* Golang is smart enough here to handle cases like i == 0 and i == n, etc.
