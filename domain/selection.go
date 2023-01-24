package domain

import (
	"context"
	"github.com/shopspring/decimal"
)

type OutcomeState int

func (o OutcomeState) ToString() string {
	switch o {
	case Unsettled:
		return "Unsettled"
	case Void:
		return "Void"
	case Lose:
		return "Lose"
	case Place:
		return "Place"
	case Win:
		return "Win"
	}
	return ""
}

const (
	Unsettled OutcomeState = iota
	Void
	Lose
	Place
	Win
)

type Selection struct {
	Id             int
	Name           string
	SelectedEvent  Event
	SelectedMarket Market
	Price          decimal.Decimal
	IsActive       bool `json:"is_active"`
	Outcome        OutcomeState
}
type SelectionUseCase interface {
	CreateSelection(ctx context.Context, selection Selection, marketId, eventId int) (Selection, error)
	UpdateSelection(ctx context.Context, selection Selection, selectionId int) (Selection, error)
	DeleteSelection(ctx context.Context, selectionId int) error
	DeactivateSelection(ctx context.Context, selectionId int) error
	ActivateSelection(ctx context.Context, selectionId int) error
}
