package header

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"log"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/grpc-ecosystem/go-grpc-middleware/logging/zap/ctxzap"
	"github.com/shivamk2406/Practice/pkg/constants"
	"go.uber.org/zap"
	"google.golang.org/grpc/metadata"
)

type JWTClaims struct {
	TenantID         string `json:"tenant_id,omitempty"`
	TenantRecordID   int64  `json:"tenant_record_id,omitempty"`
	CustomerID       string `json:"customer_id,omitempty"`
	DeviceID         string `json:"device_id,omitempty"`
	TenantCustomerID string `json:"tenant_customer_id,omitempty"`
}

const (
	jwtClaimsHeaderKey string = "x-jwt-claims"
)

func ProcessClaims() func(ctx context.Context, md metadata.MD) context.Context {
	return func(ctx context.Context, md metadata.MD) context.Context {
		// TODO: need to move logger deps out of pkg.
		logger := ctxzap.Extract(ctx)

		if len(md[jwtClaimsHeaderKey]) < 1 {
			logger.Panic("missing jwt claims in request")
		}

		data := md[jwtClaimsHeaderKey][0]

		var payload JWTClaims

		b, err := base64.StdEncoding.DecodeString(data)
		if err != nil {
			logger.Panic("failed to decode jwt payload from request due to", zap.Error(err))
		}

		err = json.Unmarshal(b, &payload)
		if err != nil {
			logger.Panic("failed to build jwt payload due to %s", zap.Error(err))
		}

		return context.WithValue(ctx, constants.JWTClaimsKey, payload)
	}
}

// HeaderFunc reads the incoming grpc metadata and sets it into context.
func HeadersFunc(ctx context.Context) context.Context {
	defer func() {
		if e := recover(); e != nil {
			log.Printf("recovered from panic due to %v", e)
		}
	}()

	headers := make(map[string][]byte)
	md, ok := ctx.Value(constants.GrpcMetadataKey).(metadata.MD)
	if !ok {
		return context.WithValue(ctx, constants.KafkaHeadersKey, headers)
	}

	for _, h := range ForwardHeaders {
		if len(md[h]) > 0 {
			headers[h] = []byte(md[h][0])
		}
	}
	return context.WithValue(ctx, constants.KafkaHeadersKey, headers)
}

func (c *JWTClaims) Validate() error {
	a := validation.ValidateStruct(c,
		validation.Field(&c.CustomerID, validation.Required),
		validation.Field(&c.DeviceID, validation.Required),
		validation.Field(&c.TenantID, validation.Required),
		validation.Field(&c.TenantRecordID, validation.Required),
		validation.Field(&c.TenantCustomerID, validation.Required))

	return a
}
