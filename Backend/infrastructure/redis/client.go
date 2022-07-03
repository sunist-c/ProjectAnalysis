package redis

import (
	"github.com/go-redis/redis"
	"time"
)

type Client struct {
	client *redis.Client
}

func (c *Client) Connect(cfg Config) (err error) {
	c.client = redis.NewClient(&redis.Options{
		Addr:         cfg.Address,
		Password:     cfg.Password,
		DB:           cfg.toDatabaseIndex(),
		MaxRetries:   cfg.toMaxRetry(),
		DialTimeout:  cfg.toTimeout(),
		ReadTimeout:  cfg.toTimeout(),
		WriteTimeout: cfg.toTimeout(),
		PoolTimeout:  cfg.toTimeout(),
		IdleTimeout:  cfg.toTimeout(),
	})

	_, err = c.client.Ping().Result()
	return
}

// Write set a key-value pair into redis
func (c Client) Write(key, value string, timeout time.Duration) (err error) {
	_, err = c.client.Set(key, value, timeout).Result()
	return
}

// Read get a key-value pair in redis
func (c Client) Read(key string) (result string, err error) {
	return c.client.Get(key).Result()
}

// Remove delete a key-value pair in redis
func (c Client) Remove(key string) (err error) {
	_, err = c.client.Del(key).Result()
	return
}
