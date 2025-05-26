package search

import "context"

type Searcher interface {
	Search(ctx context.Context, query string) ([]Result, error)
}

type Result struct {
	FileUrl string
	Repo    string
}
