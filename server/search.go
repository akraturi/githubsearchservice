package server

import (
	"context"
	v1 "githubsearchservice/gen/github.com/akraturi/githubsearchservice/pkg/pb/v1"
)

func (server *Server) Search(ctx context.Context, r *v1.SearchRequest) (*v1.SearchResponse, error) {
	var results []*v1.Result
	results = append(results, &v1.Result{
		FileUrl: "dummy file url",
		Repo:    "dummy repo",
	})

	return &v1.SearchResponse{Results: results}, nil
}
