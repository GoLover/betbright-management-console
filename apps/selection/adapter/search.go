package adapter

import (
	"betbright-management-console/domain"
	"context"
)

type SearchAdapter interface {
	Search(ctx context.Context, index, query string) (domain.SearchResult, error)
}
