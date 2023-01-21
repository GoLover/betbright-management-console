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
	GetSportById(id int) (Sport, error)
	UpdateSport(sport Sport, sportSlug string) (Sport, error)
	ChangeActivationSport(sportSlug string, active bool) error
	CreateEvent(event Event, sportSlug string) (Event, error)
	GetEventById(id int) (Event, error)
	UpdateEvent(event Event, eventSlug, sportSlug string) (Event, error)
	ChangeActivationEvent(eventSlug string, active bool) (Event, error)
	GetEventsBySportId(sportId int) ([]Event, error)
	CreateMarket(market Market, eventSlug string) (Market, error)
	GetMarketById(id int) (Market, error)
	UpdateMarket(market Market, marketId int, eventSlug string) (Market, error)
	ChangeActivationMarket(marketId int, active bool) (Market, error)
	GetMarketsByEventId(eventId int) ([]Market, error)
	CreateSelection(selection Selection, marketId, eventId int) (Selection, error)
	GetSelectionById(id int) (Selection, error)
	UpdateSelection(selection Selection, selectionId, marketId, eventId int) (Selection, error)
	ChangeActivationSelection(selectionId int, active bool) (Selection, error)
	GetSelectionByMarketId(marketId int) ([]Selection, error)
}
