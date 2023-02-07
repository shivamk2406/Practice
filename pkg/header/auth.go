package header

import (
	"context"
	"strings"

	"github.com/grpc-ecosystem/go-grpc-middleware/logging/zap/ctxzap"
	"github.com/shivamk2406/Practice/pkg/constants"
	"google.golang.org/grpc/metadata"
)

type AuthClaims struct {
	CustomerID string
}

const (
	authHeaderKey string = "authorization"
)

func ProcessAuth() func(ctx context.Context, md metadata.MD) context.Context {
	return func(ctx context.Context, md metadata.MD) context.Context {
		// TODO: need to move logger deps out of pkg.
		logger := ctxzap.Extract(ctx)

		for _, v := range md[authHeaderKey] {
			logger.Info("inside authorization map")
			logger.Info(v)
		}
		if len(md[authHeaderKey]) < 1 {
			logger.Panic("missing auth header in request")
		}

		data, _ := extractTokenFromAuthHeader(md[authHeaderKey][0])

		return context.WithValue(ctx, constants.AuthKey, data)
	}
}

func extractTokenFromAuthHeader(val string) (token string, ok bool) {
	authHeaderParts := strings.Split(val, " ")
	if len(authHeaderParts) != 2 || !strings.EqualFold(authHeaderParts[0], "bearer") {
		return "", false
	}

	return authHeaderParts[1], true
}
