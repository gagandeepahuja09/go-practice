package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"path"
)

// this error is returned when the avatar instance is unable to provide an avatar URL.
// note: this will be created only once and everytime its pointer will be used, which
// is pretty inexpensive.
var ErrNoAvatarURL = errors.New("chat: Unable to get an avatar URL")

// represents types capable of representing user profile picture.
type Avatar interface {
	// ErrNoAvatarURL is returned, if the object is unable to get any URL from the
	// client.
	GetAvatarURL(ChatUser) (string, error)
}

type AuthAvatar struct{}

var UseAuthAvatar AuthAvatar

// this doesn't have a very nice line of sight as the return is buried inside
// refactor it.
func (AuthAvatar) GetAvatarURL(u ChatUser) (string, error) {
	url := u.AvatarURL()
	if len(url) == 0 {
		return "", ErrNoAvatarURL
	}
	return url, nil
}

type GravatarAvatar struct{}

var UseGravatar GravatarAvatar

func (GravatarAvatar) GetAvatarURL(u ChatUser) (string, error) {
	return fmt.Sprintf("//www.gravatar.com/avatar/%x", u.UniqueID()), nil
}

type FileSystemAvatar struct{}

var UseFileSystemAvatar FileSystemAvatar

func (FileSystemAvatar) GetAvatarURL(u ChatUser) (string, error) {
	if files, err := ioutil.ReadDir("avatars"); err == nil {
		for _, file := range files {
			if file.IsDir() {
				continue
			}
			if match, _ := path.Match(u.UniqueID()+"*", file.Name()); match {
				return "/avatars" + file.Name(), nil
			}
		}
	}
	return "", ErrNoAvatarURL
}
