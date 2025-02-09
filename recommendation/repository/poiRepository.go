package repository

import (
	"bytes"
	"context"
	"encoding/json"
	"recommendation/dto"
	"recommendation/infrastructure"
	"sync"

	"github.com/elastic/go-elasticsearch/v8"
	log "github.com/sirupsen/logrus"
)

var (
	poiRepositoryInit     sync.Once
	poiRepositoryInstance *poiRepository
)

type (
	PoiRepository interface {
		SearchPoiByTitle(c context.Context, titleVector []float64, scriptQuery infrastructure.ResponseScriptQuery, indexName string) []dto.PoiEntity
	}
	poiRepository struct {
		client *elasticsearch.Client
	}
)

func NewPoiRepository(client *elasticsearch.Client) PoiRepository {
	poiRepositoryInit.Do(func() {
		poiRepositoryInstance = &poiRepository{
			client: client,
		}
	})
	return poiRepositoryInstance
}

func (p poiRepository) SearchPoiByTitle(c context.Context, titleVector []float64, script infrastructure.ResponseScriptQuery, indexName string) []dto.PoiEntity {

	script.Script.Query.ScriptScore.Script.Params.QueryVector = titleVector
	script.Script.Source = []string{"title", "address", "location"}
	reqBody := marshalJson(script.Script)
	res, err := poiRepositoryInstance.client.Search(
		poiRepositoryInstance.client.Search.WithContext(c),
		poiRepositoryInstance.client.Search.WithBody(bytes.NewReader(reqBody)),
		poiRepositoryInstance.client.Search.WithIndex(indexName),
		poiRepositoryInstance.client.Search.WithTrackTotalHits(true),
	)
	if err != nil {
		log.Printf("Failed to execute search : %v", err)
		return nil
	}
	defer res.Body.Close()
	if res.IsError() {
		log.Errorf("Error: %+v", res.String())
	}

	// 응답 분문 파싱 및 로깅
	var response map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&response); err != nil {
		log.Errorf("Failed to decode response : %v", err)
		return nil
	}

	// 결과 처리
	hits := response["hits"].(map[string]interface{})["hits"].([]interface{})
	var poiList []dto.PoiEntity
	for _, hit := range hits {
		source := hit.(map[string]interface{})["_source"]
		sourceJSON, err := json.Marshal(source)
		if err != nil {
			log.Errorf("Failed to marshal source : %v", err)
			continue
		}
		var esResp dto.PoiEntity
		if err := json.Unmarshal(sourceJSON, &esResp); err != nil {
			log.Errorf("Failed to unmarshal source : %v", err)
			continue
		}
		poiList = append(poiList, esResp)
	}
	return poiList
}

func marshalJson(v any) []byte {
	marshal, err := json.Marshal(v)
	if err != nil {
		log.Errorf("error : %v", err.Error())
	}
	return marshal
}
