package elasticsearch

import (
	"betbright-management-console/domain"
	"betbright-management-console/infra/elastic"
	"fmt"
	"strconv"
)

type SearchRepository struct {
	elasticDb *elastic.Client
}

func (s *SearchRepository) Update(index string, data map[string]interface{}) error {
	id := strconv.Itoa(int(data["id"].(int64)))
	return s.elasticDb.Update(index, id, data)
}

func (s *SearchRepository) SearchIndex(query string) (domain.SearchResult, error) {
	s.elasticDb.Search(`*`, query)
	return domain.SearchResult{}, nil
}

func (s *SearchRepository) Delete(index, key string) error {
	err := s.elasticDb.Delete(index, index, key)
	return err
}

func (s *SearchRepository) Insert(index string, data map[string]interface{}) error {
	br := &elastic.BulkRequest{
		Action: elastic.ActionCreate,
		Index:  index,
		Data:   data,
		ID:     strconv.Itoa(int(data["id"].(int64))),
	}
	bulk, err := s.elasticDb.IndexBulk(index, []*elastic.BulkRequest{br})
	if err != nil {
		return err
	}
	fmt.Printf(`%#v`, bulk)
	return nil
}

func New(elasticDb *elastic.Client) *SearchRepository {
	return &SearchRepository{elasticDb: elasticDb}
}
