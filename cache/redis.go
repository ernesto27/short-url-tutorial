package cache

import (
	"crypto/tls"
	"fmt"

	"github.com/go-redis/redis"
)

type Redis struct {
	redis *redis.Client
}

func NewRedis(host string, password string, port string) *Redis {
	tlsConfig := &tls.Config{MinVersion: tls.VersionTLS12}
	if password == "" {
		tlsConfig = nil
	}

	rc := redis.NewClient(&redis.Options{
		Addr:      host + ":" + port,
		Password:  password,
		DB:        0,
		TLSConfig: tlsConfig,
	})

	err := rc.Ping().Err()
	if err != nil {
		fmt.Printf("failed to connect with redis instance at %s - %v", host, err)
	}

	return &Redis{
		redis: rc,
	}
}

func (r *Redis) Set(key string, value interface{}) error {
	return r.redis.Set(key, value, 0).Err()
}

func (r *Redis) Get(key string) (string, error) {
	return r.redis.Get(key).Result()
}
