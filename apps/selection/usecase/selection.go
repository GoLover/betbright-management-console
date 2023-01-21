package usecase

import (
	"betbright-management-console/domain"
	"context"
	"fmt"
)

type SelectionUseCase struct {
	observers []domain.Observer
	r         domain.SportRepository
}

func (s *SelectionUseCase) Register(observer domain.Observer) {
	if s.observers == nil {
		s.observers = make([]domain.Observer, 0)
	}
	s.observers = append(s.observers, observer)
}

func (s *SelectionUseCase) Notify(ctx context.Context) {
	for _, k := range s.observers {
		k.Update(ctx)
	}
}

func (s *SelectionUseCase) CreateSelection(ctx context.Context, selection domain.Selection, marketId, eventId int) (domain.Selection, error) {
	selection.IsActive = true
	return s.r.CreateSelection(selection, marketId, eventId)
}

func (s *SelectionUseCase) UpdateSelection(ctx context.Context, selection domain.Selection, selectionId int) (domain.Selection, error) {
	//TODO implement me
	panic("implement me")
}

func (s *SelectionUseCase) DeleteSelection(ctx context.Context, selectionId int) error {
	//TODO implement me
	panic("implement me")
}

func (s *SelectionUseCase) DeactivateSelection(ctx context.Context, id int) error {
	selection, err := s.r.ChangeActivationSelection(id, false)
	if err != nil {
		return err
	}
	s.Notify(context.WithValue(ctx, `marketId`, selection.SelectedMarket.Id))
	fmt.Println(`selection deactivated successfully`)
	return err
}
func (s *SelectionUseCase) ActivateSelection(ctx context.Context, id int) error {
	_, err := s.r.ChangeActivationSelection(id, true)
	return err
}

func New(r domain.SportRepository) *SelectionUseCase {
	return &SelectionUseCase{r: r}
}
