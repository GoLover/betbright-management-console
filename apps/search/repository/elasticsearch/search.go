package elasticsearch

import (
	"betbright-management-console/domain"
	"betbright-management-console/infra/elastic"
	"encoding/json"
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

func (s *SearchRepository) SearchIndex(index, query string) (domain.SearchResult, error) {
	searchResult := domain.SearchResult{
		Sports:    make([]domain.Sport, 0),
		Events:    make([]domain.Event, 0),
		Markets:   make([]domain.Market, 0),
		Selection: make([]domain.Selection, 0),
	}
	elasticResp, err := s.elasticDb.Search(index, query)
	if err != nil {
		return searchResult, err
	}
	if elasticResp.Hits.Total.Value == 0 {
		return searchResult, domain.ErrSearchHasNoResult
	}
	for _, k := range elasticResp.Hits.Hits {
		searchData, _ := json.Marshal(k.Source)
		switch k.Index {
		case "sports":
			sport := domain.Sport{}
			json.Unmarshal(searchData, &sport)
			searchResult.Sports = append(searchResult.Sports, sport)
		case "markets":
			market := domain.Market{}
			json.Unmarshal(searchData, &market)
			searchResult.Markets = append(searchResult.Markets, market)
		case "events":
			event := domain.Event{}
			json.Unmarshal(searchData, &event)
			searchResult.Events = append(searchResult.Events, event)
		case "selections":
			selection := domain.Selection{}
			json.Unmarshal(searchData, &selection)
			searchResult.Selection = append(searchResult.Selection, selection)
		}
	}
	return searchResult, nil
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
