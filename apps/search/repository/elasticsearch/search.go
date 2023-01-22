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

func (s *SearchRepository) SearchIndex(index, query string) (domain.SearchResult, error) {
	//TODO implement me
	panic("implement me")
}

func (s *SearchRepository) Delete(key string) error {
	//TODO implement me
	panic("implement me")
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
