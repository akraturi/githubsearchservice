package search

import (
	"context"
	"fmt"
	"github.com/go-resty/resty/v2"
)

type GithubSearcher struct {
	githubApiClient *resty.Client
}

func NewGithubSearcher(githubApiClient *resty.Client) *GithubSearcher {
	return &GithubSearcher{
		githubApiClient: githubApiClient,
	}
}

type repository struct {
	FullName string `json:"full_name"`
}

type item struct {
	HtmlUrl    string     `json:"html_url"`
	Repository repository `json:"repository"`
}

type githubSearchApiResponse struct {
	Items []item `json:"items"`
}

func (gs *GithubSearcher) Search(ctx context.Context, query string) ([]Result, error) {
	var searchResp githubSearchApiResponse

	resp, err := gs.githubApiClient.R().
		SetContext(ctx).
		SetQueryParam("q", query).
		SetResult(&searchResp).
		Get("search/code")

	if err != nil {
		return nil, fmt.Errorf("api call to github failed: %v", err)
	}

	if resp == nil {
		return nil, fmt.Errorf("github search api returned empty response")
	}

	if resp.IsError() {
		return nil, fmt.Errorf("api call to github failed with status: %v", resp.Status())
	}

	var searchResults []Result
	for _, item := range searchResp.Items {
		searchResults = append(searchResults, Result{
			FileUrl: item.HtmlUrl,
			Repo:    item.Repository.FullName,
		})
	}

	return searchResults, nil
}
