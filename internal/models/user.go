package models

import (
	"github.com/google/uuid"
)

type User struct {
	Profile   UserProfile
	TimeTable UserTimeTable
}

type UserProfile struct {
	ID      uuid.UUID
	Name    string
	Surname string
	Age     int
	Sex     string
}

type UserTimeTable struct {
	Events []Event
}
