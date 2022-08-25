package kafka

type Scheduler interface {
	Topic() string
}

type Topic string

func (t Topic) Topic() string {
	return string(t)
}
