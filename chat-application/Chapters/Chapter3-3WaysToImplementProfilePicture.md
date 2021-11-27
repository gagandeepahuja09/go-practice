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