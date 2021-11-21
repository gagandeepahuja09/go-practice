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