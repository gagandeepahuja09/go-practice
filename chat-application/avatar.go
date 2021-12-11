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
	GetAvatarURL(c *client) (string, error)
}

type AuthAvatar struct{}

var UseAuthAvatar AuthAvatar

// this doesn't have a very nice line of sight as the return is buried inside
// refactor it.
func (AuthAvatar) GetAvatarURL(c *client) (string, error) {
	if url, ok := c.userData["avatar_url"]; ok {
		if urlString, ok := url.(string); ok {
			return urlString, nil
		}
	}
	return "", ErrNoAvatarURL
}

type GravatarAvatar struct{}

var UseGravatar GravatarAvatar

func (GravatarAvatar) GetAvatarURL(c *client) (string, error) {
	if userId, ok := c.userData["userid"]; ok {
		if userIdStr, ok := userId.(string); ok {
			return fmt.Sprintf("//www.gravatar.com/avatar/%x", userIdStr), nil
		}
	}
	return "", ErrNoAvatarURL
}

type FileSystemAvatar struct{}

var UseFileSystemAvatar FileSystemAvatar

func (FileSystemAvatar) GetAvatarURL(c *client) (string, error) {
	if userId, ok := c.userData["userid"]; ok {
		if userIdStr, ok := userId.(string); ok {
			files, err := ioutil.ReadDir("avatars")
			if err != nil {
				return "", ErrNoAvatarURL
			}
			for _, file := range files {
				if file.IsDir() {
					continue
				}
				if match, _ := path.Match(userIdStr+"*", file.Name()); match {
					return "/avatars" + file.Name(), nil
				}
			}
		}
	}
	return "", ErrNoAvatarURL
}
