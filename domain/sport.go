package domain

import (
	"context"
)

type Sport struct {
	Id          int
	Name        string
	DisplayName string
	Slug        string
	Order       int
	IsActive    bool
	Events      []Event
}

type SportUseCase interface {
	CreateSport(ctx context.Context, sport Sport) (Sport, error)
	UpdateSport(ctx context.Context, sport Sport, sportSlug string) (Sport, error)
	DeactivateSport(ctx context.Context, slug string) error
	ActivateSport(ctx context.Context, slug string) error
}

type SportRepository interface {
	CreateSport(sport Sport) (Sport, error)
	UpdateSport(sport Sport, sportSlug string) (Sport, error)
	CreateEvent(event Event, sportSlug string) (Event, error)
	UpdateEvent(event Event, eventSlug, sportSlug string) (Event, error)
	CreateMarket(market Market, eventSlug string) (Market, error)
	UpdateMarket(market Market, marketId int, eventSlug string) (Market, error)
	CreateSelection(selection Selection, marketId, eventId int) (Selection, error)
	UpdateSelection(selection Selection, selectionId, marketId, eventId int) (Selection, error)
	ChangeActivationSport(sportSlug string, active bool) error
	ChangeActivationEvent(eventSlug string, active bool) error
	ChangeActivationMarket(marketId int, active bool) error
	ChangeActivationSelection(selectionId int, active bool) error
}
