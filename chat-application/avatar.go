package main

import "errors"

// this error is returned when the avatar instance is unable to provide an avatar URL.
// note: this will be created only once and everytime its pointer will be used, which
// is pretty inexpensive.
var ErrNoAvatarURL = errors.New("chat: Unable to get an avatar URL.")

// represents types capable of representing user profile picture.
type Avatar interface {
	// ErrNoAvatarURL is returned, if the object is unable to get any URL from the
	// client.
	GetAvatarURL(c *client) (string, error)
}
