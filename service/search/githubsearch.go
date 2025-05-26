package search

import (
	"context"
	"fmt"
	"github.com/go-resty/resty/v2"
	"os"
	"time"
)

type GithubSearcher struct {
	githubApiClient *resty.Client
}

func NewGithubSearcher() *GithubSearcher {
	githubApiClient := resty.New().
		SetBaseURL("https://api.github.com").
		SetTimeout(5*time.Second).
		SetRetryCount(3).
		AddRetryAfterErrorCondition().
		SetHeader("Content-Type", "application/json").
		SetHeader("Authorization",
			fmt.Sprintf("token %v", os.Getenv("GITHUB_TOKEN")))

	return &GithubSearcher{
		githubApiClient: githubApiClient,
	}
}

type githubSearchApiResponse struct {
	Items []struct {
		HtmlUrl    string `json:"html_url"`
		Repository struct {
			FullName string `json:"full_name"`
		} `json:"repository"`
	} `json:"items"`
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
