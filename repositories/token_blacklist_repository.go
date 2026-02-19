package repositories

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

type TokenBlacklistRepository struct {
	Redis *redis.Client
}

func NewTokenBlacklistRepository(rdb *redis.Client) *TokenBlacklistRepository {
	return &TokenBlacklistRepository{
		Redis: rdb,
	}
}

func (r *TokenBlacklistRepository) AddToken(token string, expiresAt time.Time) error {
	ctx := context.Background()
	duration := time.Until(expiresAt)
	return r.Redis.Set(ctx, token, "blacklisted", duration).Err()
}

func (r *TokenBlacklistRepository) IsTokenBlacklisted(token string) (bool, error) {
	ctx := context.Background()
	val, err := r.Redis.Exists(ctx, token).Result()
	if err != nil {
		return false, err
	}
	return val > 0, nil
}
