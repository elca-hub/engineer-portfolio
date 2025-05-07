package database

import (
	"context"
	"devport/domain/repo/nosql"
	"devport/infra/database/redis/redis_repo"
	"fmt"

	"github.com/redis/go-redis/v9"
)

type RedisHandler struct {
	client *redis.Client
}

type RedisRepositoryConfig struct {
	client *redis.Client
}

func NewRedisRepositoryConfig(client *redis.Client) *RedisRepositoryConfig {
	return &RedisRepositoryConfig{client: client}
}

func (c *RedisRepositoryConfig) UserRepository() nosql.UserRepository {
	return redis_repo.NewRedisUserRepository(c.client)
}

func NewRedisHandler(c *RedisConfig) (NoSQLInter, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", c.host, c.port),
		Password: c.password,
		PoolSize: c.poolSize,
	})

	_, err := rdb.Ping(context.Background()).Result()
	if err != nil {
		return nil, err
	}

	return NewRedisRepositoryConfig(rdb), nil
}
