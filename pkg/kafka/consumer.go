package kafka

import (
	"context"
	"fmt"

	"github.com/grpc-ecosystem/go-grpc-middleware/logging/zap/ctxzap"
	"github.com/segmentio/kafka-go"
)

type ConsumerConfig struct {
	BootstrapServers []string `json:"bootstrap_servers,omitempty"`
	Topic            string   `json:"topic,omitempty"`
	Group            string   `json:"group,omitempty"`
}

type Consumer struct {
	reader    *kafka.Reader
	config    *kafka.ReaderConfig
	onConsume ConsumeFunc
}

type ConsumerOption func(*Consumer)

type ConsumeFunc func(c context.Context, b []byte) error

func WithOnConsume(f ConsumeFunc) ConsumerOption {
	return func(c *Consumer) {
		c.onConsume = f
	}
}

func NewConsumer(_ context.Context, cfg ConsumerConfig, opts ...ConsumerOption) (*Consumer, error) {
	consumer := &Consumer{
		config: &kafka.ReaderConfig{
			Brokers: cfg.BootstrapServers,
			GroupID: cfg.Group,
			Topic:   cfg.Topic,
		},
		onConsume: func(c context.Context, b []byte) error { return nil },
	}

	for _, opt := range opts {
		opt(consumer)
	}

	return consumer, nil
}

// Start initialized the consumer. The msg will be processed using func onConsume().
// The context passed to this func will need to invoke pkg.ProcessFromMetadata()
// to extract all required headers.
func (c Consumer) Start(ctx context.Context) {
	logger := ctxzap.Extract(ctx).Sugar()

	defer func() {
		if e := recover(); e != nil {
			logger.Errorf("recovered from panic due to %v", e)
		}
	}()

	c.reader = kafka.NewReader(*c.config)

	defer c.reader.Close()

	fmt.Println("____________________:" + c.config.Topic)
	logger.Info("consumer: starting for topic %s", c.config.Topic)

	for {
		msg, err := c.reader.ReadMessage(context.Background())
		if err != nil {
			logger.Errorf("consumer: error %v occurred for topic %s", err, c.config.Topic)
			continue
		}

		val := msg.Value

		newCtx := ExtractHeaders(ctx, msg.Headers)

		logger.Infof("consumer: successfully consumed msg on topic:%v and data is: %v", c.config.Topic, string(val))

		err = c.onConsume(newCtx, val)
		if err != nil {
			logger.Errorf("consumer: error %v while processing msg from topic %s", err, c.config.Topic)
		}
	}
}
