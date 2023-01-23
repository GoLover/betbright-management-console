package usecase

import (
	"betbright-management-console/domain"
	"context"
	"strconv"
)

type SearchUsecase struct {
	r domain.SearchRepository
}

func (s *SearchUsecase) Search(ctx context.Context, index, query string) (domain.SearchResult, error) {
	return s.r.SearchIndex(index, query)
}

func (s *SearchUsecase) Update(ctx context.Context, index string, data map[string]interface{}) error {
	return s.r.Update(index, data)
}

func (s *SearchUsecase) Delete(ctx context.Context, index string, data map[string]interface{}) error {
	id := strconv.Itoa(int(data[`id`].(int64)))
	return s.r.Delete(index, id)
}

func (s *SearchUsecase) Index(ctx context.Context, index string, data map[string]interface{}) error {
	err := s.r.Insert(index, data)
	return err
}

func New(r domain.SearchRepository) *SearchUsecase {
	return &SearchUsecase{r: r}
}
