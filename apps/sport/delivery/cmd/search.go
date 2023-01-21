package cmd

import (
	"betbright-management-console/domain"
	"context"
)

type SearchOperator struct {
	u domain.SportUseCase
}

func (s SearchOperator) Create(ctx context.Context) {
	s.SearchAll(ctx)
}

func (s SearchOperator) Update(ctx context.Context) {
	s.SearchAll(ctx)
}

func (s SearchOperator) Delete(ctx context.Context) {
	s.SearchAll(ctx)
}

func (s SearchOperator) Search(ctx context.Context) {
	s.SearchAll(ctx)
}

func (s SearchOperator) SearchAll(ctx context.Context) {
	//TODO implement me
	panic("implement me")
}
