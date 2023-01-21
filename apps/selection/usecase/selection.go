package usecase

import (
	"betbright-management-console/domain"
	"context"
)

type SelectionUseCase struct {
	r domain.SportRepository
}

func (s SelectionUseCase) CreateSelection(ctx context.Context, selection domain.Selection, marketId, eventId int) (domain.Selection, error) {
	return s.r.CreateSelection(selection, marketId, eventId)
}

func (s SelectionUseCase) UpdateSelection(ctx context.Context, selection domain.Selection, selectionId int) (domain.Selection, error) {
	//TODO implement me
	panic("implement me")
}

func (s SelectionUseCase) DeleteSelection(ctx context.Context, selectionId int) error {
	//TODO implement me
	panic("implement me")
}

func (s SelectionUseCase) DeactivateSelection(ctx context.Context, id int) error {
	//TODO implement me
	panic("implement me")
}

func New(r domain.SportRepository) *SelectionUseCase {
	return &SelectionUseCase{r: r}
}
