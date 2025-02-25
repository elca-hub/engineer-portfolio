package redis_repo

import (
	"context"
	"github.com/redis/go-redis/v9"
	"devport/domain/model"
	"time"
)

type RedisUserRepository struct {
	client *redis.Client
}

func NewRedisUserRepository(client *redis.Client) *RedisUserRepository {
	return &RedisUserRepository{
		client: client,
	}
}

func (r *RedisUserRepository) StartSession(email *model.Email) (string, error) {
	token := model.NewUUID("").ID()

	if err := r.client.Set(context.Background(), token, email.Email(), time.Hour).Err(); err != nil {
		return "", err
	}

	return token, nil
}

func (r *RedisUserRepository) GetSession(token string) (*model.Email, error) {
	token, err := r.client.Get(context.Background(), token).Result()

	if err != nil {
		return nil, err
	}

	email, err := model.NewEmail(token)

	if err != nil {
		return nil, err
	}

	return email, nil
}

func (r *RedisUserRepository) DeleteSession(token string) error {
	if err := r.client.Del(context.Background(), token).Err(); err != nil {
		return err
	}

	return nil
}
