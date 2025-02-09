package service

import (
	"context"
	"recommendation/dto"
	"recommendation/infrastructure"
	"recommendation/repository"
	"sync"
)

var (
	searchServiceInit     sync.Once
	searchServiceInstance *searchService
)

type (
	SearchService interface {
		SearchTitle(ctx context.Context, title string) []dto.PoiEntity
	}

	searchService struct {
		repository      repository.PoiRepository
		embeddingServer infrastructure.ModelServiceCaller
	}
)

/*
*
SearachService 생성 메서드
*/
func NewSearchService(repository repository.PoiRepository,
	recommendationRepository repository.RecommendationRepsoitory,
	embeddingServer infrastructure.ModelServiceCaller) SearchService {

	searchServiceInit.Do(func() {
		searchServiceInstance = &searchService{
			repository:      repository,
			embeddingServer: embeddingServer,
		}
	})
	return searchServiceInstance
}

/*
*
 */
func (s *searchService) SearchTitle(c context.Context, title string) []dto.PoiEntity {
	embedResponse := s.getVector(c, title)
	scriptQuery := s.getQueryByIntent(title)
	return s.repository.SearchPoiByTitle(c, embedResponse.Vector, scriptQuery, "vector_poi")
}

func (s *searchService) getVector(c context.Context, title string) infrastructure.ResponseEmbedQuery {
	return s.embeddingServer.CallEmbedQuery(
		&infrastructure.RequestEmbedQuery{
			Text: title,
		})
}

func (s *searchService) getQueryByIntent(query string) infrastructure.ResponseScriptQuery {
	return s.embeddingServer.CallQueryByIntent(
		&infrastructure.RequestScriptQuery{
			Text: query,
		})
}
