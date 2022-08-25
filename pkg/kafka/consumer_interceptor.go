package kafka

import (
	"context"

	"github.com/segmentio/kafka-go"
	"github.com/shivamk2406/Practice/pkg/constants"
	"google.golang.org/grpc/metadata"
)

func ExtractHeaders(ctx context.Context, headers []kafka.Header) context.Context {
	md := metadata.MD{}

	for _, h := range headers {
		md[h.Key] = []string{string(h.Value)}
	}

	return context.WithValue(ctx, constants.GrpcMetadataKey, md)
}
