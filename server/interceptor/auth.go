package interceptor

import (
	"context"
	"githubsearchservice/pkg"
	"log"
	"strings"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

func Auth() grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (interface{}, error) {
		log.Println("incoming request: ", info.FullMethod, req)

		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			return nil, status.Errorf(codes.Unauthenticated, "metadata is not provided")
		}

		values := md["github_api_token"]
		if len(values) == 0 {
			return nil, status.Errorf(codes.Unauthenticated, "github api token is not provided")
		}

		accessToken := values[0]
		if len(strings.TrimSpace(accessToken)) == 0 {
			return nil, status.Errorf(codes.Unauthenticated, "empty github api token")
		}

		return handler(pkg.GetContextWithGithubAPIToken(ctx, accessToken), req)
	}
}
