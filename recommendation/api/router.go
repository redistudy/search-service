package api

import (
	v1 "recommendation/api/v1"
	"recommendation/infrastructure"
	"recommendation/repository"
	"recommendation/service"
	"recommendation/setting"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

func SetRouters(r *gin.Engine, cfg *setting.Configuration, client *elasticsearch.Client, redisClient *redis.Client) {
	poi := createPoiDomain(cfg, client, redisClient)
	apiv1 := r.Group("/api/v1")
	{
		// Create artile with tags
		// Get aritle detail by id
		apiv1.POST("/poi", poi.SearchPoi)
		// Update poi detail by id
		apiv1.PUT("/poi/:id", poi.UpdatePoi)
		// Delete poi by id
		apiv1.DELETE("/poi/:id", poi.DeletePoi)

		// Add other router if necessary
	}
}

// @Summary POI 전체 도메인 생성 내부 메서드
func createPoiDomain(cfg *setting.Configuration, client *elasticsearch.Client, redisClient *redis.Client) v1.PoiController {
	modelServer := infrastructure.NewModelServerCaller()
	poiRepository := repository.NewPoiRepository(client)
	// TODO : 추천 서비스 추가
	recommendationRepository := repository.NewRecommendationRepository(redisClient)
	service := service.NewSearchService(poiRepository, recommendationRepository, modelServer)
	poiRouter := v1.NewPoiRouter(service, cfg)
	return poiRouter
}
