package domain

import (
	"context"
	"fmt"
	"github.com/pterm/pterm"
	"strconv"
	"strings"
)

type Table struct {
	Schema string
	Name   string

	Columns   []TableColumn
	PKColumns []int

	UnsignedColumns []int
}
type TableColumn struct {
	Name       string
	Type       int
	Collation  string
	RawType    string
	IsAuto     bool
	IsUnsigned bool
	IsVirtual  bool
	IsStored   bool
	EnumValues []string
	SetValues  []string
	FixedSize  uint
	MaxSize    uint
}

// Rule is the rule for how to sync data from MySQL to ES.
// If you want to sync MySQL data into elasticsearch, you must set a rule to let use know how to do it.
// The mapping rule may thi: schema + table <-> index + document type.
// schema and table is for MySQL, index and document type is for Elasticsearch.
type Rule struct {
	Schema string   `toml:"schema"`
	Table  string   `toml:"table"`
	Index  string   `toml:"index"`
	Type   string   `toml:"type"`
	Parent string   `toml:"parent"`
	ID     []string `toml:"id"`

	// Default, a MySQL table field name is mapped to Elasticsearch field name.
	// Sometimes, you want to use different name, e.g, the MySQL file name is title,
	// but in Elasticsearch, you want to name it my_title.
	FieldMapping map[string]string `toml:"field"`

	// MySQL table information
	TableInfo *Table

	//only MySQL fields in filter will be synced , default sync all fields
	Filter []string `toml:"filter"`

	// Elasticsearch pipeline
	// To pre-process documents before indexing
	Pipeline string `toml:"pipeline"`
}

func newDefaultRule(schema string, table string) *Rule {
	r := new(Rule)

	r.Schema = schema
	r.Table = table

	lowerTable := strings.ToLower(table)
	r.Index = lowerTable
	r.Type = lowerTable

	r.FieldMapping = make(map[string]string)

	return r
}

func (r *Rule) prepare() error {
	if r.FieldMapping == nil {
		r.FieldMapping = make(map[string]string)
	}

	if len(r.Index) == 0 {
		r.Index = r.Table
	}

	if len(r.Type) == 0 {
		r.Type = r.Index
	}

	// ES must use a lower-case Type
	// Here we also use for Index
	r.Index = strings.ToLower(r.Index)
	r.Type = strings.ToLower(r.Type)

	return nil
}

// CheckFilter checkers whether the field needs to be filtered.
func (r *Rule) CheckFilter(field string) bool {
	if r.Filter == nil {
		return true
	}

	for _, f := range r.Filter {
		if f == field {
			return true
		}
	}
	return false
}

type SearchQuery struct {
}

type SearchResult struct {
	Sports    []Sport
	Events    []Event
	Markets   []Market
	Selection []Selection
}

func (sr SearchResult) PrettyPrint() {
	if len(sr.Sports) > 0 {
		ptd := make([][]string, 0)
		ptd = append(ptd, []string{"ID", "Name", "DisplayName", "Slug", "Order", "isActive"})
		for _, k := range sr.Sports {
			ptd = append(ptd, []string{strconv.Itoa(k.Id), k.Name, k.DisplayName, k.Slug, strconv.Itoa(k.Order), fmt.Sprint(k.IsActive)})
		}
		pterm.DefaultTable.WithHasHeader().WithData(ptd).Render()
	}
	if len(sr.Events) > 0 {
		ptd := make([][]string, 0)
		ptd = append(ptd, []string{"ID", "Name", "Slug", "Type", "Status", "SportID", "isActive"})
		for _, k := range sr.Events {
			ptd = append(ptd, []string{strconv.Itoa(k.Id), k.Name, k.Slug, k.EType.ToString(), k.Status.ToString(),
				strconv.Itoa(k.SportId), fmt.Sprint(k.IsActive)})
		}
		pterm.DefaultTable.WithHasHeader().WithData(ptd).Render()
	}
	if len(sr.Markets) > 0 {
		ptd := make([][]string, 0)
		ptd = append(ptd, []string{"ID", "Name", "DisplayName", "EventId", "Schema", "Column", "Order", "isActive"})
		for _, k := range sr.Markets {
			ptd = append(ptd, []string{strconv.Itoa(k.Id), k.Name, k.DisplayName, strconv.Itoa(k.EventId),
				strconv.Itoa(k.Schema), strconv.Itoa(k.Columns), strconv.Itoa(k.Order), fmt.Sprint(k.IsActive)})
		}
		pterm.DefaultTable.WithHasHeader().WithData(ptd).Render()
	}
	if len(sr.Selection) > 0 {
		ptd := make([][]string, 0)
		ptd = append(ptd, []string{"ID", "Name", "EventId", "MarketId", "Outcome", "Price", "isActive"})
		for _, k := range sr.Selection {
			ptd = append(ptd, []string{strconv.Itoa(k.Id), k.Name, strconv.Itoa(k.SelectedEvent.Id), strconv.Itoa(k.SelectedMarket.Id), k.Outcome.ToString(), k.Price.String(), fmt.Sprint(k.IsActive)})
		}
		pterm.DefaultTable.WithHasHeader().WithData(ptd).Render()
	}
}

type SearchUsecase interface {
	Search(ctx context.Context, index, query string) (SearchResult, error)
	Index(ctx context.Context, index string, data map[string]interface{}) error
	Delete(ctx context.Context, index string, data map[string]interface{}) error
	Update(ctx context.Context, index string, data map[string]interface{}) error
}

type SearchRepository interface {
	Insert(index string, data map[string]interface{}) error
	SearchIndex(index, query string) (SearchResult, error)
	Delete(index, id string) error
	Update(index string, data map[string]interface{}) error
}
