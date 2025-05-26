package search

import "context"

type Searcher interface {
	Search(ctx context.Context, query string) ([]Result, error)
}

type Result struct {
	FileURL string
	Repo    string
}
