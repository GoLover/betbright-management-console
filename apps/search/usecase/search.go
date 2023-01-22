package usecase

import (
	"betbright-management-console/domain"
	"context"
)

type SearchUsecase struct {
	r domain.SearchRepository
}

func (s *SearchUsecase) Search(ctx context.Context, query string) {
	//TODO implement me
	panic("implement me")
}

func (s *SearchUsecase) DeleteIndex(ctx context.Context, index string) error {
	//TODO implement me
	panic("implement me")
}

func (s *SearchUsecase) Index(ctx context.Context, index string, data map[string]interface{}) error {
	s.r.Insert(index, data)
	return nil
}

func New(r domain.SearchRepository) *SearchUsecase {
	return &SearchUsecase{r: r}
}
