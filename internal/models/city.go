package models

import (
	"github.com/google/uuid"
)

type City struct {
	Id     uuid.UUID
	Name   string
	Guides []Guide
}
