package pubsub

import (
	"context"

	"github.com/go-redis/redis/v8"
)

type RedisSubscriptionChannel struct {
	client  *redis.Client
	channel string
}

func NewRedisSubsciptionChannel(client *redis.Client, channel string) *RedisSubscriptionChannel {
	return &RedisSubscriptionChannel{
		client:  client,
		channel: channel,
	}
}

func (s *RedisSubscriptionChannel) Subscribe(handler func(message string)) context.CancelFunc {
	ctx, cancel := context.WithCancel(context.Background())
	pubsub := s.client.Subscribe(ctx, s.channel)
	go func() {
		msgChan := pubsub.Channel()
		for {
			select {
			case msg := <-msgChan:
				handler(msg.Payload)
			}
		}
	}()
	return cancel
}
