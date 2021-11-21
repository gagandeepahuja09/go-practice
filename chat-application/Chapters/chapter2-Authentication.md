What we will learn:

* Use decorator pattern to wrap http.Handler types to add additional functionality 
 to handlers.
* Serve HTTP endpoints with dynamic paths.
* Use Gomniauth open source project to access auth services.
* Get and set cookies using the http package.
* Send and receive JSON data over a websocket.
* Give different types of data to templates.
* Work with channels of our own type.

* authHandler will just redirect to another handler if all conditions are met.

******************************

oAuth2
* Open authentication & authorization standard designed to allow resource owners to 
 give clients delegated access to private data (such as wall posts or tweets) via an
 exchange token handshake.
* Even if we don't want to access the private data, if gives a great option to sign in
 using their existing credentials without exposing them to the 3rd party site.

oAuth2 flow for user's point of view:
* User selects the provider.
* Redirected to provider's website where they are asked to give permission to the 
 client app.
* User accpets the permission
* User is redirected to the client app with a request code.
* The client app sends the grant code to the provider, who sends back an auth token.
* The client app uses the access token to make authorized requests to the provider.


Tell the authentication providers about our app
* There is generally a dashbaord / console for it.
* In order to identify the client application, we need to create a client ID and secret.
* We will also be asked about the redirect URI, the endpoint to which the user will 
be redirected after signing in.
* The callback will be another action on our loginHandler, so the redirection URL will 
be like: localhost:8080/auth/callback/google.

* add client id, secret for the google oauth from console to the application
* redirect to the redirect url using gomniAuth.Provider and GetBeginAuthURL and redirecting to it by setting the response writer header location and setting the status code of temporary redirect
* While redirecting we would also be provided with the grant code by the provider: 
http://localhost:8080/auth/callback/google?code=4%2F0AX4XfWiFOXr5opMmy_8fpIhwzW5rsSl4NRUVxyTOsoWw7JE5xUC11-PlfVKF-5U0SKoz9Q&scope=email+profile+https%3A%2F%2Fwww.googleapis.com%2Fauth%2Fuserinfo.email+https%3A%2F%2Fwww.googleapis.com%2Fauth%2Fuserinfo.profile+openid&authuser=0&prompt=consent#
* We don't have to worry of what to do with this code as gomniAuth will process it for us.
* This code will be exchanged by the auth provider for a token that allows us to
access the private user data. For added security, this step happens behind the scenes,
from server to server rather than within the browser.


* Implementing our callback handler.
* The callback handler will be the redirect uri.
* Extract the creds from code=query param.
* Extract the user from the creds. The above 2 steps will be done by a 3rd party oAuth proivder like  gomniauth.
* After extract the user details, we will set the cookie which will contain the user name, base 64 encoded.


* Presenting the user data
* Update the serveHTTP of the template handler so that we can inject some data into it.


* Augmenting messages with additional data
* Additional data to be added in message: message sender, timestamp.
* create a new message struct.
* In client.go, we will replace readMessage, writeMessage with readJSON, writeJSON.
* Replace all structs with using message struct instead of []byte.
* In the handler of /room, we create a new client. While creating the new 
client we should set the userData field also in it.
* This will be later used in message.Name.
* While reading from the frontend in the read method implemented in client.go, we 
make sure that we set the name and timestamp.
* While sending and receiving from the frontend, we make sure that we pass the required JSON(stringified) and show the appropriate name and message.
