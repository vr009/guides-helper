package models

import (
	"github.com/google/uuid"
)

type Guide struct {
	Profile   GuideProfile
	Expertise GuideExpertise
	TimeTable GuideTimeTable
}

type GuideProfile struct {
	ID      uuid.UUID
	Name    string
	Surname string
	Age     int
	Sex     string
}

type GuideExpertise struct {
	Cities []City
}

type GuideTimeTable struct {
	Events []Event
	//TODO to look how to work with time table entities
}
