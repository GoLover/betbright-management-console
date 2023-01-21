package usecase

import (
	"betbright-management-console/domain"
	"context"
)

type MarketUseCase struct {
	r domain.SportRepository
}

func (s MarketUseCase) CreateMarket(ctx context.Context, market domain.Market, eventSlug string) (domain.Market, error) {
	return s.r.CreateMarket(market, eventSlug)
}

func (s MarketUseCase) UpdateMarket(ctx context.Context, market domain.Market, marketId int) (domain.Market, error) {
	//TODO implement me
	panic("implement me")
}

func (s MarketUseCase) DeleteMarket(ctx context.Context, marketId int) error {
	//TODO implement me
	panic("implement me")
}

func (s MarketUseCase) DeactivateMarket(ctx context.Context, marketId int) error {
	//TODO implement me
	panic("implement me")
}


func New(r domain.SportRepository) *MarketUseCase {
	return &MarketUseCase{r: r}
}
