package domain

import (
	"context"
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
