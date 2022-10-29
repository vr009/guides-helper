package models

import (
	"time"
)

type EventType string

const (
	Excursion  = "Excursion"
	Trip       = "Trip"
	Exhibition = "Exhibition"
)

type Event struct {
	Type  EventType
	Guide Guide
	City  City
	Users []User
	When  time.Time
}
