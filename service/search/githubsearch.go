package search

import (
	"context"
	"fmt"
	"githubsearchservice/pkg"

	"github.com/go-resty/resty/v2"
)

type GithubSearcher struct {
	githubAPIClient *resty.Client
}

func NewGithubSearcher(githubAPIClient *resty.Client) *GithubSearcher {
	return &GithubSearcher{
		githubAPIClient: githubAPIClient,
	}
}

type repository struct {
	FullName string `json:"full_name"`
}

type item struct {
	HTMLURL    string     `json:"html_url"`
	Repository repository `json:"repository"`
}

type githubSearchAPIResponse struct {
	Items []item `json:"items"`
}

func (gs *GithubSearcher) Search(ctx context.Context, query string) ([]Result, error) {
	var searchResp githubSearchAPIResponse

	githubAPIToken := pkg.GetGithubAPIToken(ctx)

	resp, err := gs.githubAPIClient.R().
		SetHeader("Authorization",
			fmt.Sprintf("token %v", githubAPIToken)).
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
			FileURL: item.HTMLURL,
			Repo:    item.Repository.FullName,
		})
	}

	return searchResults, nil
}
