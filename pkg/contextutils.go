package pkg

import "context"

type contextKey int

const githubAPITokenKey contextKey = iota

func GetContextWithGithubAPIToken(ctx context.Context, token string) context.Context {
	return context.WithValue(ctx, githubAPITokenKey, token)
}

func GetValueFromContext[T any](ctx context.Context, key contextKey) T {
	if ctx != nil {
		if value, ok := ctx.Value(key).(T); ok {
			return value
		}
	}
	return *new(T)
}

func GetGithubAPIToken(ctx context.Context) string {
	return GetValueFromContext[string](ctx, githubAPITokenKey)
}
