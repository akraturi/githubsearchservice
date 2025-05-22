package githubservice

import (
	"context"
	"fmt"
	"github.com/go-resty/resty/v2"
	"os"
	"time"
)

type GithubService struct {
	githubApiClient *resty.Client
}

func New() GithubService {
	githubApiClient := resty.New().
		SetBaseURL("https://api.github.com").
		SetTimeout(5*time.Second).
		SetRetryCount(3).
		AddRetryAfterErrorCondition().
		SetHeader("Content-Type", "application/json").
		SetHeader("Authorization",
			fmt.Sprintf("token %v", os.Getenv("GITHUB_TOKEN")))

	return GithubService{
		githubApiClient: githubApiClient,
	}
}

type GithubSearchApiResponse struct {
	Items []struct {
		HtmlUrl    string `json:"html_url"`
		Repository struct {
			FullName string `json:"full_name"`
		} `json:"repository"`
	} `json:"items"`
}

func (g *GithubService) Search(ctx context.Context, query string) (GithubSearchApiResponse, error) {
	searchResp := GithubSearchApiResponse{}

	resp, err := g.githubApiClient.R().
		SetContext(ctx).
		SetQueryParam("q", query).
		SetResult(&searchResp).
		Get("search/code")

	if err != nil {
		return searchResp, fmt.Errorf("api call to github failed: %v", err)
	}

	if resp == nil {
		return searchResp, fmt.Errorf("github search api returned empty response")
	}

	if resp.IsError() {
		return searchResp, fmt.Errorf("api call to github failed with status: %v", resp.Status())
	}

	return searchResp, nil
}
