package kafka

import (
	"context"

	"github.com/grpc-ecosystem/go-grpc-middleware/logging/zap/ctxzap"
	"github.com/pkg/errors"
	"github.com/segmentio/kafka-go"
	"github.com/shivamk2406/Practice/pkg/header"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

type ProducerConfig struct {
	BootstrapServers []string `json:"bootstrap_servers,omitempty"`
	Topic            string   `json:"topic,omitempty"`
}

type Producer struct {
	marshaller      protojson.MarshalOptions
	protoMarshaller proto.MarshalOptions
	protojson       bool
	writer          *kafka.Writer
	headersFunc     func(ctx context.Context) []kafka.Header
}

type ProducerOption func(*Producer)

func WithHeadersFunc(f func(ctx context.Context) []kafka.Header) ProducerOption {
	return func(p *Producer) {
		p.headersFunc = f
	}
}

func WithProtoJSONMarshalOpts(m protojson.MarshalOptions) ProducerOption {
	return func(p *Producer) {
		p.marshaller = m
	}
}

func WithProtoMarshaller(m proto.MarshalOptions) ProducerOption {
	return func(p *Producer) {
		p.protoMarshaller = m
	}
}

func WithProto() ProducerOption {
	return func(p *Producer) {
		p.protojson = false
	}
}

func NewProducer(_ context.Context, cfg ProducerConfig, opts ...ProducerOption) (*Producer, error) {
	producer := &Producer{
		writer: &kafka.Writer{
			Addr:         kafka.TCP(cfg.BootstrapServers...),
			Topic:        cfg.Topic,
			Async:        false,
			RequiredAcks: kafka.RequireAll,
		},
		headersFunc:     func(ctx context.Context) []kafka.Header { return []kafka.Header{} },
		marshaller:      protojson.MarshalOptions{},
		protoMarshaller: proto.MarshalOptions{},
		protojson:       true,
	}

	for _, opt := range opts {
		opt(producer)
	}

	return producer, nil
}

func (p *Producer) Produce(ctx context.Context, pb proto.Message, opts ...ProducerOption) error {
	logger := ctxzap.Extract(ctx).Sugar()
	var val []byte
	var err error

	if p.protojson {
		val, err = p.marshaller.Marshal(pb)
	} else {
		val, err = p.protoMarshaller.Marshal(pb)
	}
	if err != nil {
		return errors.WithMessage(err, "producer: failed to marshal message")
	}

	for _, opt := range opts {
		opt(p)
	}

	msg := kafka.Message{
		Value: val,
		// TODO: add trace id
		Headers: p.headersFunc(header.HeadersFunc(ctx)),
	}

	err = p.writer.WriteMessages(ctx, msg)
	if err != nil {
		return errors.WithMessagef(err, "producer: failed publish msg to topic %s", p.writer.Topic)
	}

	logger.Info("produced for topic: ", p.writer.Topic)
	return nil
}
