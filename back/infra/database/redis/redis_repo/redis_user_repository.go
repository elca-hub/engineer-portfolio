package redis_repo

import (
	"context"
	"devport/domain/model"
	"time"

	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
)

type RedisUserRepository struct {
	client *redis.Client
}

func NewRedisUserRepository(client *redis.Client) *RedisUserRepository {
	return &RedisUserRepository{
		client: client,
	}
}

func (r *RedisUserRepository) IsExistSession(token string) (bool, error) {

	if err := r.client.Exists(context.Background(), token).Err(); err != nil {
		return false, err
	}

	return true, nil
}

func (r *RedisUserRepository) StartSession(email *model.Email) (string, error) {
	tokenRaw, err := uuid.NewUUID()

	if err != nil {
		return "", err
	}

	token := tokenRaw.String()

	if err := r.client.Set(context.Background(), token, email.Email(), time.Hour*24).Err(); err != nil {
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
