package domain

import "context"

type Market struct {
	Id          int
	Name        string
	DisplayName string `json:"display_name"`
	Order       int
	IsActive    bool `json:"is_active"`
	Schema      int
	Columns     int
	EventId     int
}
type MarketUseCase interface {
	CreateMarket(ctx context.Context, market Market, eventSlug string) (Market, error)
	UpdateMarket(ctx context.Context, market Market, marketId int, eventSlug string) (Market, error)
	DeleteMarket(ctx context.Context, marketId int) error
	DeactivateMarket(ctx context.Context, marketId int) error
	ActivateMarket(ctx context.Context, marketId int) error
}
