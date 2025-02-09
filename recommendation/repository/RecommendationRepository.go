package repository

import (
	"context"
	"sync"

	"github.com/redis/go-redis/v9"
)

var (
	RecommendationRepsoitoryInit     sync.Once
	RecommendationRepositoryInstance *recommendationRepository
	recommendationNamespace          = "recommendation"
)

type (
	RecommendationRepsoitory interface {
		RecommendaFromUserSearchLogFeature(c context.Context, userId string)
	}
	recommendationRepository struct {
		redisClient *redis.Client
	}
)

func NewRecommendationRepository(redisClient *redis.Client) RecommendationRepsoitory {
	RecommendationRepsoitoryInit.Do(func() {
		RecommendationRepositoryInstance = &recommendationRepository{
			redisClient: redisClient,
		}
	})
	return RecommendationRepositoryInstance
}

func (recommendation recommendationRepository) RecommendaFromUserSearchLogFeature(c context.Context, userId string) {
	// TODO : Redis로부터 UserFeature 가져오는 로직 구현
}
