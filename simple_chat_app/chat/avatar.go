package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"path"
)

var ErrNoAvatarURL = errors.New("chat: Unable to get an avatar URL.")

type Avatar interface {
	GetAvatarURL(c *client) (string, error)
}

type AuthAvatar struct {
}

type GravatarAvatar struct {
}

type FileSystemAvatar struct {
}

var UseAuthAvatar AuthAvatar
var UseGravatar GravatarAvatar
var UseFileSystemAvatar FileSystemAvatar

func (AuthAvatar) GetAvatarURL(c *client) (string, error) {
	if url, ok := c.userData["avatar_url"]; ok {
		if urlStr, ok := url.(string); ok {
			return urlStr, nil
		}
	}
	return "", ErrNoAvatarURL
}

func (GravatarAvatar) GetAvatarURL(c *client) (string, error) {
	if userId, ok := c.userData["userId"]; ok {
		if userIdStr, ok := userId.(string); ok {
			return fmt.Sprintf("//www.gravatar.com/avatar/%s", userIdStr), nil
		}
	}
	return "", ErrNoAvatarURL
}
func (FileSystemAvatar) GetAvatarURL(c *client) (string, error) {
	if avatarUrl, ok := c.userData["avatar_url"]; ok {
		if avatarUrlStr, ok := avatarUrl.(string); ok {
			return avatarUrlStr, nil
		}
	}
	if userId, ok := c.userData["userId"]; ok {
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
					return fmt.Sprintf("/avatars/%s", file.Name()), nil
				}
			}
		}
	}
	return "", ErrNoAvatarURL
}
