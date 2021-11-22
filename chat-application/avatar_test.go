package main

import "testing"

// if error is returned, it should always be ErrNoAvatarURL.
// happy flow, set the user data for client.
func TestAuthAvatar(t *testing.T) {
	// we are making use of Go's zero initialization capabilities.
	// since there is no state for our object, we won't waste any memory in
	// initializing it.
	var authAvatar AuthAvatar
	client := new(client)
	_, err := authAvatar.GetAvatarURL(client)
	if err != ErrNoAvatarURL {
		t.Error("authAvatar.GetAvatarURL should return ErrNoAvatarURL when no avatar URL present")
	}

	testURL := "https://test-url.com"
	client.userData = map[string]interface{}{
		"avatar_url": testURL,
	}
	url, err := authAvatar.GetAvatarURL(client)
	if err != nil {
		t.Error("AuthAvatar.GetAvatarURL should return no error when value present")
	}
	if url != testURL {
		t.Error("AuthAvatar.GetAvatarURL should return correct URL")
	}
}

func TestGravatarAvatar(t *testing.T) {
	var gravatarAvatar GravatarAvatar
	client := new(client)
	client.userData = map[string]interface{}{
		"email": "MyEmailAddress@example.com",
	}

	url, err := gravatarAvatar.GetAvatarURL(client)
	if err != nil {
		t.Error("gravatarAvatar.GetAvatarURL should not return an error")
	}
	if url != "//www.gravatar.com/avatar/0bc83cb571cd1c50ba6f3e8a78ef1346" {
		t.Errorf("gravatarAvatar.GetAvatarURL wrongly returned %s", url)
	}
}
