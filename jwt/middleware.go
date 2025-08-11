package jwt

import (
	"context"
	"errors"
	"strings"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type contextKey string

var (
	ErrNoAuthHeader = errors.New("authorization header missing")
	ErrInvalidAuth  = errors.New("invalid authorization header")
	userIDKey       = contextKey("user_id")
	roleKey         = contextKey("role")
)

// gRPC middleware для проверки JWT токена
func UnaryJWTInterceptor(jm JWTManager) grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (interface{}, error) {
		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			return nil, ErrNoAuthHeader
		}
		authHeader := md["authorization"]
		if len(authHeader) == 0 {
			return nil, ErrNoAuthHeader
		}
		token := strings.TrimPrefix(authHeader[0], "Bearer ")
		claims, err := jm.Verify(token)
		if err != nil {
			return nil, ErrInvalidAuth
		}
		ctx = context.WithValue(ctx, userIDKey, claims.UserID)
		ctx = context.WithValue(ctx, roleKey, claims.Role)
		return handler(ctx, req)
	}
}

func UserIDFromContext(ctx context.Context) (int64, bool) {
	id, ok := ctx.Value(userIDKey).(int64)
	return id, ok
}

func RoleFromContext(ctx context.Context) (string, bool) {
	role, ok := ctx.Value(roleKey).(string)
	return role, ok
}
