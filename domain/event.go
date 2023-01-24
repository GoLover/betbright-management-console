package domain

import "context"

type EventType int

func (e EventType) ToString() string {
	switch e {
	case ETPrePlay:
		return `Pre Play`
	case ETInPlay:
		return `In Play`
	}
	return ``
}

type EventStatus int

func (e EventStatus) ToString() string {
	switch e {
	case ESPrePlay:
		return `Pre Play`
	case ESInPlay:
		return `In Play`
	case ESEnded:
		return `Ended`
	}
	return ``
}

const (
	_ EventType = iota
	ETPrePlay
	ETInPlay
)

const (
	_ EventStatus = iota
	ESPrePlay
	ESInPlay
	ESEnded
)

type Event struct {
	Id       int
	Name     string
	EType    EventType
	Status   EventStatus
	Slug     string
	IsActive bool `json:"is_active"`
	SportId  int  `json:"sport_id"`
	Markets  []Market
}

type EventUseCase interface {
	CreateEvent(ctx context.Context, event Event, sportSlug string) (Event, error)
	UpdateEvent(ctx context.Context, event Event, eventSlug, sportSlug string) (Event, error)
	DeleteEvent(ctx context.Context, slug string) error
	DeactivateEvent(ctx context.Context, slug string) error
	ActivateEvent(ctx context.Context, slug string) error
}
