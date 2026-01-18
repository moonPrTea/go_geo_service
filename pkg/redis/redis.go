package redis

import (
	"context"
	"github.com/redis/go-redis/v9"
)

type Queue struct {
	Client *redis.Client
}

func Init(addr string) *Queue {
	rdb := redis.NewClient(&redis.Options{Addr: addr})
	return &Queue{Client: rdb}
}

func (q *Queue) Enqueue(payload string) {
	q.Client.RPush(context.Background(), "webhooks", payload)
}

func (q *Queue) Push(payload string) error {
    context := context.Background()
    err := q.Client.RPush(context, "webhooks", payload).Err()
    if err != nil {
        return err
    }
    return nil
}
