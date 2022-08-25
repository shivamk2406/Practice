package kafka

import (
	"log"

	"github.com/segmentio/kafka-go"
	"github.com/shivamk2406/Practice/pkg/constants"
	"golang.org/x/net/context"
)

func InjectHeaders(ctx context.Context) []kafka.Header {
	defer func() {
		if e := recover(); e != nil {
			log.Printf("recovered from panic due to %v", e)
		}
	}()

	headers := make([]kafka.Header, 0)

	claims, ok := ctx.Value(constants.KafkaHeadersKey).(map[string][]byte)
	if !ok {
		log.Printf("producer: invalid headers found %v: %v", string(constants.KafkaHeadersKey), ctx.Value(constants.KafkaHeadersKey))
		return headers
	}

	for k, v := range claims {
		headers = append(headers, kafka.Header{
			Key:   k,
			Value: v,
		})
	}
	return headers
}
