package usecase

import (
	"betbright-management-console/domain"
	"context"
	"errors"
	"fmt"
)

type MarketUseCase struct {
	observers         []domain.Observer
	subjectsToObserve []domain.Observee
	r                 domain.SportRepository
}

func (s *MarketUseCase) Update(ctx context.Context) {
	marketId := ctx.Value(`marketId`).(int)
	_, err := s.r.GetSelectionByMarketId(marketId)
	if errors.Is(err, domain.ErrRepoRecordNotFound) {
		err = s.DeactivateMarket(ctx, marketId)
		if err != nil {
			fmt.Println(fmt.Errorf(`DeactiveMarket-UpdateSignal %w`, err))
			return
		}
		return
	}
	if err != nil {
		fmt.Println(fmt.Errorf(`Market-UpdateSignal %w`, err))
	}

}

func (s *MarketUseCase) Register(observer domain.Observer) {
	if s.observers == nil {
		s.observers = make([]domain.Observer, 0)
	}
	s.observers = append(s.observers, observer)
}

func (s *MarketUseCase) Notify(ctx context.Context) {
	for _, k := range s.observers {
		k.Update(ctx)
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
	market, err := s.r.ChangeActivationMarket(marketId, false)
	if err != nil {
		return err
	}
	s.Notify(context.WithValue(ctx, `eventId`, market.EventId))
	fmt.Println(`market deactivated successfully`)
	return err
}
func (s *MarketUseCase) ActivateMarket(ctx context.Context, marketId int) error {
	_, err := s.r.ChangeActivationMarket(marketId, true)
	//if err != nil {
	//	return err
	//}
	//s.Notify(context.WithValue(ctx, `eventId`, market.EventId))
	return err
}

func (s *MarketUseCase) GetMarketsByEventId(eventId int) ([]domain.Market, error) {
	return s.r.GetMarketsByEventId(eventId)
}
func (s *MarketUseCase) BindObserveLately(subjectsToObserve []domain.Observee) {
	for _, k := range subjectsToObserve {
		k.Register(s)
	}
}

func New(r domain.SportRepository, subjectsToObserve []domain.Observee) *MarketUseCase {
	mu := &MarketUseCase{r: r}
	for _, k := range subjectsToObserve {
		k.Register(mu)
	}
	return mu
}
