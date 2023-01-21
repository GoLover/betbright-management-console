package domain

import (
	"context"
	"github.com/shopspring/decimal"
)

type OutcomeState int

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
	IsActive       bool
	Outcome        OutcomeState
}
type SelectionUseCase interface {
	CreateSelection(ctx context.Context, selection Selection, marketId, eventId int) (Selection, error)
	UpdateSelection(ctx context.Context, selection Selection, selectionId int) (Selection, error)
	DeleteSelection(ctx context.Context, selectionId int) error
	DeactivateSelection(ctx context.Context, selectionId int) error
	ActivateSelection(ctx context.Context, selectionId int) error
}
