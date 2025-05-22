## Github Search Service

A Grpc service built on top of the GitHub API to perform queries for the
provided search phrase and allows for optional filtering down to the user level. It returns the
file URL and the repo it was found in.

API used to perform the GitHub search:
https://docs.github.com/en/rest/reference/search

### API Spec
````protobuf
service GithubSearchService {
  rpc Search(SearchRequest) returns (SearchResponse);
}

message SearchRequest {
  string search_term = 1;
  string user = 2;
}

message SearchResponse {
  repeated Result results = 1;
}

message Result {
  string file_url = 1;
  string repo = 2;
}
````