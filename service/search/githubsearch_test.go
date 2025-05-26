package search

import (
	"context"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/suite"
	"net/http"
	"testing"
)

type githubSearcherTestSuite struct {
	suite.Suite
	githubSearcher  *GithubSearcher
	githubApiClient *resty.Client
}

func (s *githubSearcherTestSuite) SetupSuite() {
	s.githubApiClient = resty.New().SetBaseURL("https://api.github.com")
	s.githubSearcher = NewGithubSearcher(s.githubApiClient)
	httpmock.ActivateNonDefault(s.githubApiClient.GetClient())
}

func (s *githubSearcherTestSuite) TearDownTest() {
	httpmock.DeactivateAndReset()
}

func (s *githubSearcherTestSuite) TestSearch_Success() {
	expectedGithubApiResponse := githubSearchApiResponse{
		Items: []item{
			{
				HtmlUrl: "https://github.com/user/repo1/blob/master/file1.go",
				Repository: repository{
					FullName: "user/repo1",
				},
			},
			{
				HtmlUrl: "https://github.com/user/repo2/blob/master/file2.go",
				Repository: repository{
					FullName: "user/repo2",
				},
			},
		},
	}

	expectedResult := []Result{
		{FileUrl: "https://github.com/user/repo1/blob/master/file1.go", Repo: "user/repo1"},
		{FileUrl: "https://github.com/user/repo2/blob/master/file2.go", Repo: "user/repo2"},
	}

	mockResponder, err := httpmock.NewJsonResponder(200, expectedGithubApiResponse)
	httpmock.RegisterResponder("GET", fmt.Sprintf("https://api.github.com/search/code"),
		mockResponder)
	results, err := s.githubSearcher.Search(context.Background(), "test")

	s.Require().NoError(err)
	s.Require().Equal(len(expectedResult), len(results))

	for i, result := range results {
		s.Require().Equal(result.FileUrl, expectedGithubApiResponse.Items[i].HtmlUrl)
		s.Require().Equal(result.Repo, expectedGithubApiResponse.Items[i].Repository.FullName)
	}
}

func (s *githubSearcherTestSuite) TestSearch_Failure_On_ApiFailure() {
	mockResponder := httpmock.ResponderFromResponse(&http.Response{
		Status:        "500",
		StatusCode:    500,
		ContentLength: -1,
	})

	httpmock.RegisterResponder("GET", fmt.Sprintf("https://api.github.com/search/code"),
		mockResponder)
	_, err := s.githubSearcher.Search(context.Background(), "test")

	s.Require().Error(err)
	s.Require().ErrorContains(err, "500")
}

func TestGithubSearcher_Search(t *testing.T) {
	suite.Run(t, new(githubSearcherTestSuite))
}
