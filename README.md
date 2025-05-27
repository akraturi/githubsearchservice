# Github Search Service
A Grpc service built on top of the GitHub API to perform queries for the
provided search phrase and allows for optional filtering down to the user level. It returns the
file URL and the repo it was found in.

API used to perform the GitHub search:
https://docs.github.com/en/rest/reference/search

## API Spec
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

## Prerequisites
### Linux:
- Make and Git: `sudo apt-get install git make buf`
   
### Mac:
- Make and Git: `brew install git make bufbuild/buf/buf`

## Setup:
- Clone the repository: `git clone https://github.com/akraturi/githubsearchservice.git`
- Install dev dependencies
   ```sh
   make dev-setup
   ```
- Generate protobuf code (if modified): `make gen`

## Run
### Server
- Run
  ```sh
   go run main.go server
  ```

### Client
- While server is running same application can be used as grpc client
- Github api token is required to be supplied
- To search only a term run
  ```sh
   go run main.go client --term=github --github_api_token=<your-api-token-here>
  ```
- To search optionally only for a user
  ```sh
   go run main.go client --term=github --user=akraturi --github_api_token=<your-api-token-here>
  ```
## Running Tests And Linting
- Run tests `make test`
- Run lint `make lint`

## Troubleshooting
- Server does not start: Ensure port 50051 is free currently it runs on this predefined port 
- Github API errors: Ensure client is passed a valid and active API token

### Future Enhancements
- Add more test cases
- Containerize the app to remove deps to avoid env issues
- Support for pagination
- Support for better configurable logging
- Add CI pipeline