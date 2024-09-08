package redis

import (
	"time"

	"github.com/go-redis/redis"
	"github.com/oogway93/golangArchitecture/internal/repository"
)

type Redis struct {
	client *redis.Client
}

// New creates a new instance of Redis
func New() (repository.CacheRepository, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	_, err := client.Ping().Result()
	if err != nil {
		return nil, err
	}

	return &Redis{client}, nil
}

// Set stores the value in the redis database
func (r *Redis) Set(key string, value []byte, ttl time.Duration) error {
	return r.client.Set(key, value, ttl).Err()
}

// Get retrieves the value from the redis database
func (r *Redis) Get(key string) ([]byte, error) {
	res, err := r.client.Get(key).Result()
	bytes := []byte(res)
	return bytes, err
}

// Delete removes the value from the redis database
func (r *Redis) Delete(key string) error {
	return r.client.Del(key).Err()
}

// DeleteByPrefix removes the value from the redis database with the given prefix
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
