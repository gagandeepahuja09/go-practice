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
* Rather than hardcoding * everywhere, we have used a constant(otherWord) for it.
* We have used strings.replace to replace otherWord.

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


Synonyms
* Unlike sprinkle and domainify, synonyms will print more than one result. Not an issue as our applications have the ability to handle more than one lines of input.
* Using environment variable as configuration

Using environment variables for configuration
* API key is a sensitive piece of info. If we store it in const, we won't be able to share our code without sharing our API key.
* Additionally, we'll have to recompile our entire project if the key expires or if we want to use a different one.
* PHP uses this method: read_env_file($envDir, '.env.vault') to read the env file and $appEnvironment = env('APP_ENV') for read the flag.
* For creating environment variable, you can run the export command or modify the bashrc file ~/.bashrc. Alternatively you could create multiple env files.
* And will running the deployment, it would be an environment flag.
* We'll use encoding/json package.
* Question: defer response.Body.Close(). Why do we close the response body?
* The good thing about strongly typed here is that we describe the type we expect from JSON data beforehand. Eg. we expect that the synonyms json would contain two keys noun and verb and both of them would contain a key called syn which will be slice of strings.
* We clearly indicate this in our two structs: type synonyms struct { Noun *words,Verb *words } and type words struct { Sync [] string }. Interesting thing is that here we needed two structs for describing the JSON response in Go terms.
* Since json.NewDecoder accepts an io.Reader in arguments, we can use response.Body as it is of type io.ReadCloser.
* The Decode method of json.NewDecoder accepts an interface to which it modifies. Hence here it's addressed must be passed.
* For returning the error string, we use Errors.New or fmt.Errorf.
* log.Fatalln will write the error to the standard error stream.
* We defer the closing of response body in order to keep the memory clean.
* We can create an interface of Thesaurus, so that later other implementations can also be added like mariam-webster, dictionary.com, etc.
* We have abstracted our thesaurus implementation in a package and now we can use it in our synonyms application.
* log by default writes to the standard stream. log.fatal = log.print + os.Exit
* We can use os.GetEnv to use the env variable flag.̃



Available

* Available will connect to a WHOIS server to ask for details about the domain passed.
* Programmatically parsing its response could be an issue.
* In order to check whether a domain exists or not, we will dial up a tcp connection at the whois server.
* Once the connection opens, we will simply write the domain followed by carriage return using os.Write(conn + "rn")
* In order to iterate over the connection lines in the response, we will use bufio.NewScanner as connection is also an io.Reader.
* We add a time.sleep of 1 second to ensure rate limiting at WHOIS server end.


One Program to rule them all
* Everytime we have to run, we have to type long messy lines and pipe them.
* Here we need to build them. Overtime, the commands could change / increase or more steps could get added.
* We'll use os/exec package to run sub programs and pipe them.
* We'll keep our project in domainfinder. In that there will a lib folder which will keep build of all the subprograms. Since we don't want to copy and paste them everytime we make a change, we'll create a bash script.
* This will contain echo statements + cd commands to move to corresponding subprograms + go build commands.
* Give the execution rights to the script by running: chmod +x build.sh
* Then simply run ./build.sh to execute the bash script.
* We tie the input of first program with os.Stdin and output of last program with os.Stdout.
* When iterating through the exec.commands we ensure that the input of (i + 1)th is the output of ith command.
* For running each command, we'll use the cmd.Start method instead of Run. As run will block our program until the subprogram exists which will be no good since we'll have to run 5 subprograms at the same time.
* If the program starts successful, we'll defer a call to kill the process (cmd.Process.Kill)
* We then iterate over all commands and wait for them to finish. This is to ensure that the domainfinder doesn't exit early and kill off all the processes too soon.