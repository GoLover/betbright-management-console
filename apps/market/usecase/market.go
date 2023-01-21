package usecase

import (
	"betbright-management-console/domain"
	"context"
)

type MarketUseCase struct {
	observers         []domain.Observer
	subjectsToObserve []domain.Observee
	r                 domain.SportRepository
}

func (s *MarketUseCase) Update() {
	//TODO implement me
	panic("implement me")
}

func (s *MarketUseCase) Register(observer domain.Observer) {
	if s.observers == nil {
		s.observers = make([]domain.Observer, 0)
	}
	s.observers = append(s.observers, observer)
}

func (s *MarketUseCase) Notify() {
	for _, k := range s.observers {
		k.Update()
	}
}

func (s *MarketUseCase) CreateMarket(ctx context.Context, market domain.Market, eventSlug string) (domain.Market, error) {
	market.IsActive = true
	return s.r.CreateMarket(market, eventSlug)
}

func (s *MarketUseCase) UpdateMarket(ctx context.Context, market domain.Market, marketId int, eventSlug string) (domain.Market, error) {
	return s.r.UpdateMarket(market, marketId, eventSlug)
}

func (s *MarketUseCase) DeleteMarket(ctx context.Context, marketId int) error {
	//TODO implement me
	panic("implement me")
}

func (s *MarketUseCase) DeactivateMarket(ctx context.Context, marketId int) error {
	err := s.r.ChangeActivationMarket(marketId, false)
	if err != nil {
		return err
	}
	s.Notify()
	return err
}
func (s *MarketUseCase) ActivateMarket(ctx context.Context, marketId int) error {
	return s.r.ChangeActivationMarket(marketId, true)
}

func New(r domain.SportRepository, subjectsToObserve []domain.Observee) *MarketUseCase {
	mu := &MarketUseCase{r: r}
	for _, k := range subjectsToObserve {
		k.Register(mu)
	}
	return mu
}
