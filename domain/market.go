package domain

import "context"

type Market struct {
	Id          int
	Name        string
	DisplayName string
	Order       int
	IsActive    bool
	Schema      int
	Columns     int
}
type MarketUseCase interface {
	CreateMarket(ctx context.Context, market Market, eventSlug string) (Market, error)
	UpdateMarket(ctx context.Context, market Market, marketId int) (Market, error)
	DeleteMarket(ctx context.Context, marketId int) error
	DeactivateMarket(ctx context.Context, marketId int) error
}
