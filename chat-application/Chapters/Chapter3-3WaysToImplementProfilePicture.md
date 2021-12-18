3 Ways:
* Avatar picture provided by the auth service.
* https://en.gravatar.com/ lookup a picture on the basis of their email.
* Upload user picture.

Learnings:
* What are the good practices to get additional data from auth service, when there are no standards in place.
* When it is appropriate to build abstractions into our code.
* How Go's zero initialization pattern can save time and memory.
* How reusing an interface allows us to work with collections and individual objects in the same way as existing interfaces did.
* gravatar.
* MD5 hashing in Go.
* Upload files, storing on server.
* Serving static files through our web server.
* Using unit tests to guide the refactoring of code.
* How and when to abstract functionality from struct types into interfaces. 

***************************************************************************************

Avatars from the OAuth2 server

* To use this, we need to get the image URL from the provider, store it in the cookie for our user and send it through a websocket so that every client can render the picture alongside the corresponding message.

Getting the avatar URL
* Schema for user profile is not a part of the OAuth2 spec.
* Hence the provider does it. Different providers do things differently. Github has a avatar_url field. Google ==> picture, Fb ==> picture.url.
* Luckily gomniauth abstracts this for us. user.AvatarURL.
* We will store this in the cookie.

Logging Out
* We set the cookie and mark the maxAge as -1. After that we redirect to the /chat route which will in turn redirect to the login page.
* Max age of -1 indicates that it should be immediately deleted by the browser.
* Not all browsers are forced to delete the cookie, which is why we provide a new Value setting it to an empty string thus removing the old data that was set.
* We should update in our authHandler to handle the case where value of r.Cookie("auth") is not set. ie r.Cookie("auth").Value == "".

Abstracting the Avatar URL process
* Since we have 3 different ways of obtaining the URL, abstracting the functionality makes a lot more sense.
* Abstraction: seperating the idea of something from its implementation.
* http.Handler is a great example of this.
* The Avatar interface describes the GetAvatarURL that a type must satify in order to be able to get the avatar URL.
* We created an interface for Avatar, wrote tests for our AuthAvatar type and then
implemented it.
* code readability is important 
* const animate = true, const dontAnimate = false, move(animate), move(dontAnimate)
* is much better than move(true), move(false)

Using the implementation of authAvatar
Implementing and using Gravatar


Uploading An Avatar Picture

User Identification
* We' ll use the same approach as gravatar for user identification. MD5 hash for email.
* We' ll store this user_id in the cookie.
* This will also help in removing the ineffeciency caused by the continuous hashing. Now we need to do it only once for a user session rather than doing it for every image.

Upload Form in the UI
* The form action will point to /uploader route.
* The enctype attribute must be multipart/form-data so that the browser can transmit binary data over HTTP.
* Input of type file.
* userid is included in the form from UserData. This will tell us which user is uploading a file.

Handling the Upload
* This will receive the file, read the bytes that are streamed through the connection, and save it as a new file in the server.
* Try to write code that uses the simplest interface type you can find.
* multipart.File, multipart.Header will be returned when we try to read the req form file.
* multipart.Header will contain all metadata related to the file.
* multipart.File ==> is also an io.Reader
* we'll read it using ioutil.ReadAll()
* filename will be that of the userId along with extension taken from the header filename.

Serving the Images
* To make them accessible to the browser.
* We do this using http's built-in file server ==> http.FileServer.
* The file listing will be accessible on this route. If image, it will be displayed, else will be available for download. 
* To get a single image, we can modify the route from /avatars to /avatars/{filename.extension}

Avatar Implementation For Local Files
* We have to implement the Avatar interface for FileSystemAvatar.
* For that, let's first write tests to ensure that GetAvatarURL returns the correct required url.

Supporting different file types
* Instead of just blindly building the string, we will use ioutil.ReadDir to get a listing of the files.
* The listing will also include directories, so we will ensure that we don't read through a directory using isDir.
* We will check where each file matches the userId by a call to path.Match. Here we will use the regex pattern of userIdStr + "*" for matching.


Refactoring And Optimizing Our Code

* In our current implementation of file system avatar, we are iterating through all files in the directory and then matching for the correct file. This could become a scale problem for chatty users.
* We should be caching this during authorization at cookie level.
* Issues with doing that:
    * In order to store this url during authorization, we'll need to call getAvatarURL which has client object as one of the parameter which won't be available during the authorization.
    * One possible solution is that instead of passing client object, we'll pass all the necessary parameters.
    * Problem with this is that with each new kind of implementation, we'll need to change our interface method as well.
    * Instead we'll use an interface called chatUser which has two methods: UniqueID() and AvatarURL() in order to get the 2 properties which we care about.
* We'll create a struct which implements it and gomniauth.common.User in it as it already has the implementation of AvatarURL() and now we only need to care about UniqueID() implementation.


Changing interfaces in a test-driven way
* Before changing the implementations, we'll update the tests of each of the 3 implementations.
* Adding test of AuthAvatar. This uses gomniauth's avatar URL. Hence for tests, we'll use gomniauth's mock methods so that gomniauth's AvatarURL method returns the desired result.
* For updating the other two tests, we only need to focus on the uniqueId.

* Now we can change our all 3 implementations:
* This new implementation has not only allowed to incorporate saving of avatar url during authorization, but also helped in condensing the implementation and avoid too many null checks in all the three implementations.

Implementing our new design
* Now, we will implement our new design where we will ensure that the avatar url is set during authorization in the auth cookie.

