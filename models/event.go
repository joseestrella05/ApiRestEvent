package models

import "time"

type Event struct {
	ID          int
	Name        string
	Descripcion string
	Location    string
	DateTime    time.Time
	UserId      int
}

var events = []Event{}

func (e Event) Save() {
	// later: add it to a database
	events = append(events, e)
}
