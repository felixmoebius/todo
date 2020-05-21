package main

import "time"

type todo struct {
	Id int
	Description string
	Deadline time.Time
	Progress int
}

func formatDate(t time.Time) string {
	return t.Format("2006-01-02")
}
