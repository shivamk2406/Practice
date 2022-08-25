package header

import (
	"context"

	"google.golang.org/grpc/metadata"
)

const (
	GrpcMetadataKey contextKey = "GrpcMetadataKey"
)

func ProcessGrpcMetadata() func(ctx context.Context, md metadata.MD) context.Context {
	return func(ctx context.Context, md metadata.MD) context.Context {
		return context.WithValue(ctx, GrpcMetadataKey, md)
	}
}
