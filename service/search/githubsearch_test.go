package search

import (
	"context"
	"net/http"
	"testing"

	"github.com/go-resty/resty/v2"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/suite"
)

type githubSearcherTestSuite struct {
	suite.Suite
	githubSearcher  *GithubSearcher
	githubAPIClient *resty.Client
}

func (s *githubSearcherTestSuite) SetupSuite() {
	s.githubAPIClient = resty.New().SetBaseURL("https://api.github.com")
	s.githubSearcher = NewGithubSearcher(s.githubAPIClient)
}

func (s *githubSearcherTestSuite) SetupTest() {
	httpmock.ActivateNonDefault(s.githubAPIClient.GetClient())
}

func (s *githubSearcherTestSuite) TearDownTest() {
	httpmock.DeactivateAndReset()
}

func (s *githubSearcherTestSuite) TestSearch_Success() {
	expectedGithubAPIResponse := githubSearchAPIResponse{
		Items: []item{
			{
				HTMLURL: "https://github.com/user/repo1/blob/master/file1.go",
				Repository: repository{
					FullName: "user/repo1",
				},
			},
			{
				HTMLURL: "https://github.com/user/repo2/blob/master/file2.go",
				Repository: repository{
					FullName: "user/repo2",
				},
			},
		},
	}

	expectedResult := []Result{
		{FileURL: "https://github.com/user/repo1/blob/master/file1.go", Repo: "user/repo1"},
		{FileURL: "https://github.com/user/repo2/blob/master/file2.go", Repo: "user/repo2"},
	}

	mockResponder, _ := httpmock.NewJsonResponder(200, expectedGithubAPIResponse)
	httpmock.RegisterResponder("GET", "https://api.github.com/search/code",
		mockResponder)
	results, err := s.githubSearcher.Search(context.Background(), "test")

	s.Require().NoError(err)
	s.Require().Len(results, len(expectedResult))

	for i, result := range results {
		s.Require().Equal(result.FileURL, expectedGithubAPIResponse.Items[i].HTMLURL)
		s.Require().Equal(result.Repo, expectedGithubAPIResponse.Items[i].Repository.FullName)
	}
}

func (s *githubSearcherTestSuite) TestSearch_Failure_On_ApiFailure() {
	mockResponder := httpmock.ResponderFromResponse(&http.Response{
		Status:        "500",
		StatusCode:    500,
		ContentLength: -1,
	})

	httpmock.RegisterResponder("GET", "https://api.github.com/search/code",
		mockResponder)
	_, err := s.githubSearcher.Search(context.Background(), "test")

	s.Require().Error(err)
	s.Require().ErrorContains(err, "500")
}

func TestGithubSearcher_Search(t *testing.T) {
	suite.Run(t, new(githubSearcherTestSuite))
}
