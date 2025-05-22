package server

import (
	"context"
	"fmt"
	v1 "githubsearchservice/gen/github.com/akraturi/githubsearchservice/pkg/pb/v1"
)

func (s *Server) Search(ctx context.Context, r *v1.SearchRequest) (*v1.SearchResponse, error) {
	searchResponse := v1.SearchResponse{}

	if r.Term == "" {
		return nil, fmt.Errorf("search term cannot be empty")
	}

	searchQuery := r.GetTerm()
	if r.GetUser() != "" {
		searchQuery += " user:" + r.GetUser()
	}

	data, err := s.githubService.Search(ctx, searchQuery)
	if err != nil {
		return nil, fmt.Errorf("failed to search term: %v", err)
	}

	var results []*v1.Result
	for _, item := range data.Items {
		results = append(results, &v1.Result{
			FileUrl: item.HtmlUrl,
			Repo:    item.Repository.FullName,
		})
	}
	searchResponse.Results = results

	return &searchResponse, nil
}
