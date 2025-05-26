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
- Make and Git: `brew install git make buf`

## Setup:
- Clone the repository: `git clone https://github.com/akraturi/githubsearchservice.git`
- Install dev dependencies
   ```sh
   make dev-setup
   ```
- Generate protobuf code (if modified): `make gen`
- Update environment variables file `.env`: `GITHUB_TOKEN=<your_github_token>`

## Run
### Server
- Run
  ```sh
   go run main.go server
  ```

### Client
- While server is running same application can be used as grpc client
- To search only a term run
  ```sh
   go run main.go client --term=github
  ```
- To search optionally only for a user
  ```sh
   go run main.go client --term=github --user=akraturi
  ```

## Troubleshooting
- Server does not start: Ensure port 50051 is free currently it runs on this predefined port 
- Github API errors: Ensure environment variable `GITHUB_TOKEN` is set with a valid and active API token

### Future Enhancements
- Add more test cases
- Containerize the app to remove deps to avoid env issues
- Support better secret management
- Support for pagination