package repository

import (
	"sync"

	"github.com/redis/go-redis/v9"
)

var (
	RecommendationRepsoitoryInit     sync.Once
	RecommendationRepositoryInstance *recommendationRepository
)

type (
	RecommendationRepsoitory interface {
		RecommendaFromUserSearchLogFeature(userId string)
	}
	recommendationRepository struct {
		redisClient *redis.Client
	}
)

func NewRecommendationRepository() *recommendationRepository {
	RecommendationRepsoitoryInit.Do(func() {
		RecommendationRepositoryInstance = &recommendationRepository{}
	})
	return RecommendationRepositoryInstance
}

func (recommendation recommendationRepository) RecommendaFromUserSearchLogFeature(userId string) {
	// TODO : Redis로부터 UserFeature 가져오는 로직 구현
}
