package internal

import (
	"guides-helper/internal/models"
)

type GuidesUsecase interface {
	CreateGuide(guide models.Guide) (models.Guide, error)
	GetGuide(guide models.Guide) (models.Guide, error)
	RemoveGuide(guide models.Guide) error

	ChangeGuideProfile(guide models.Guide) (models.Guide, error)

	AddColleague(guide models.Guide, colleague models.Guide) error
	RemoveColleague(guide models.Guide, colleague models.Guide) error

	AddEventToGuide(guide models.Guide, event models.Event) error
	RemoveEventForGuide(event models.Event) error
	FixEventForGuide(event models.Event) error

	AddExpertise(guide models.Guide, expertise models.GuideExpertise) (models.Guide, error)
	ChangeExpertise(guide models.Guide, expertise models.GuideExpertise) (models.Guide, error)
	RemoveExpertise(guide models.Guide, expertise models.GuideExpertise) error
}

type GuidesRepo interface {
	InsertGuide(guide models.Guide) (models.Guide, error)
	FetchGuide(guide models.Guide) (models.Guide, error)
	DeleteGuide(guide models.Guide) error

	UpdateGuideProfile(guide models.Guide) (models.Guide, error)

	InsertColleague(guide models.Guide, colleague models.Guide) error
	DeleteColleague(guide models.Guide, colleague models.Guide) error

	InsertEventToGuide(guide models.Guide, event models.Event) error
	DeleteEventForGuide(event models.Event) error
	UpdateEventForGuide(event models.Event) error

	InsertExpertise(guide models.Guide, expertise models.GuideExpertise) (models.Guide, error)
	UpdateExpertise(guide models.Guide, expertise models.GuideExpertise) (models.Guide, error)
	DeleteExpertise(guide models.Guide, expertise models.GuideExpertise) error
}
