package pubsub

import (
	"context"

	"github.com/go-redis/redis/v8"
)

type Publisher struct {
	client *redis.Client
}

func NewRedisPublisher(client *redis.Client) *Publisher {
	return &Publisher{
		client: client,
	}
}

func (p *Publisher) Publish(channel string, message string) error {
	return p.client.Publish(context.Background(), channel, message).Err()
}
