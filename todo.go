package main

import "time"

type todo struct {
	Id int
	Description string
	Deadline time.Time
	Progress int
}
