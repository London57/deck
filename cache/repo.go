package cache

import (
	"context"
	"fmt"

	"github.com/London57/deck/entity"
	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
)


type RedisRepo struct {
	*redis.Client
}

func (r RedisRepo) GetProfile(ctx context.Context, id uuid.UUID) (entity.Profile, error) {
	var profile entity.Profile 
	err := r.HGetAll(ctx, fmt.Sprintf("profile:%s", id.String())).Scan(&profile)
	if err != nil {
		return entity.Profile{}, fmt.Errorf("failed to get profile: %w", err)
	}
	return profile, nil
}

func (r RedisRepo) SetProfile(ctx context.Context, profile entity.Profile) error {
	err := r.HSet(ctx, fmt.Sprintf("profile:%s", profile.ID.String()), profile).Err()
		if err != nil {
		return fmt.Errorf("failed to set profile: %w", err)
	}
	return nil
}