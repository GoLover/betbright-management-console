package domain

import "context"

type EventType int
type EventStatus int

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
	IsActive bool
	Markets  []Market
}

type EventUseCase interface {
	CreateEvent(ctx context.Context, event Event, sportSlug string) (Event, error)
	UpdateEvent(ctx context.Context, event Event, eventSlug, sportSlug string) (Event, error)
	DeleteEvent(ctx context.Context, slug string) error
	DeactivateEvent(ctx context.Context, slug string) error
}
