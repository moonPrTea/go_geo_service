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
	return &Queue{
		Client: rdb,
	}
}


//waiting for new event
func (q *Queue) Pop() (string, error) {
	ctx := context.Background()

	// take webhook if it exists
	result, err := q.Client.BLPop(ctx, 0, "webhooks").Result()
	if err != nil {
		return "", err
	}

	return result[1], nil
}

// push message to redis
func (q *Queue) Push(payload string) error {
    context := context.Background()
    err := q.Client.RPush(context, "webhooks", payload).Err()
	
    if err != nil {
        return err
    }
    return nil
}
