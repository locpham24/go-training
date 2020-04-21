package main

import "time"

type message struct {
	Name      string
	Message   string
	When      time.Time
	AvatarURL string
}

type outputMessage struct {
	Name      string
	Message   string
	Time      string
	AvatarURL string
}
