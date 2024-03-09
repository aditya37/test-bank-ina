package redis

import (
	"context"
	"encoding/json"
	"time"

	"github.com/adity37/task/model"
	_interface "github.com/adity37/task/repository/interface"
	cache "github.com/redis/go-redis/v9"
)

type Config struct {
	Address  string
	Password string
	DB       int
}
type redisClient struct {
	rd *cache.Client
}

func NewRedis(ctx context.Context, param Config) (_interface.RedisReaderWriter, error) {

	option := cache.Options{
		Addr:     param.Address,
		Password: param.Password,
		DB:       param.DB,
	}
	client := cache.NewClient(&option)

	// ping
	if err := client.Ping(ctx).Err(); err != nil {
		return nil, err
	}

	return &redisClient{
		rd: client,
	}, nil
}

func (nr *redisClient) Close() error {
	return nr.rd.Close()
}

func (r *redisClient) Set(ctx context.Context, key string, payload interface{}, ttl time.Duration) error {
	return r.rd.Set(ctx, key, payload, ttl).Err()
}

func (r *redisClient) GetUserSession(ctx context.Context, key string) (model.SessionPayload, error) {
	session, err := r.rd.Get(ctx, key).Result()
	if err != nil {
		return model.SessionPayload{}, err
	}
	var result model.SessionPayload
	if err := json.Unmarshal([]byte(session), &result); err != nil {
		return model.SessionPayload{}, err
	}
	return result, nil
}
