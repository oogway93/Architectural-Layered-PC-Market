package redis

import (
	"time"

	"github.com/go-redis/redis"
	"github.com/oogway93/golangArchitecture/internal/core/repository"
)

type Redis struct {
	client *redis.Client
	Config
}

type Config struct {
	Addr string
	Password string
	Expiration time.Duration
}

// New creates a new instance of Redis
func New(cfg Config) (repository.CacheRepository, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     cfg.Addr,
		Password: cfg.Password,
		DB:       0,
	})

	_, err := client.Ping().Result()
	if err != nil {
		return nil, err
	}

	return &Redis{client, cfg}, nil
}

func (r *Redis) Set(key string, value []byte) error {
	return r.client.Set(key, value, r.Config.Expiration).Err()
}

func (r *Redis) Get(key string) ([]byte, error) {
	res, err := r.client.Get(key).Result()
	bytes := []byte(res)
	return bytes, err
}

func (r *Redis) Delete(key string) error {
	return r.client.Del(key).Err()
}

func (r *Redis) DeleteByPrefix(prefix string) error {
	var cursor uint64
	var keys []string

	for {
		var err error
		keys, cursor, err = r.client.Scan(cursor, prefix, 100).Result()
		if err != nil {
			return err
		}

		for _, key := range keys {
			err := r.client.Del(key).Err()
			if err != nil {
				return err
			}
		}

		if cursor == 0 {
			break
		}
	}

	return nil
}

// Close closes the connection to the redis database
func (r *Redis) Close() error {
	return r.client.Close()
}
