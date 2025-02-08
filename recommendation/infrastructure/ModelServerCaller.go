package infrastructure

import (
	"bytes"
	json "encoding/json"
	"io/ioutil"
	"net/http"
	"sync"
)

var (
	embedQueryUrl             = "http://127.0.0.1:8000/vectorized-text"
	getQueryByIntentUrl       = "http://127.0.0.1:8000/query"
	ModelServerCallerInit     sync.Once
	modelServerCallerInstance *modelServiceCaller
)

type (
	ModelServiceCaller interface {
		CallEmbedQuery(query *RequestEmbedQuery) ResponseEmbedQuery
		CallQueryByIntent(query *RequestScriptQuery) ResponseScriptQuery
	}

	modelServiceCaller struct {
	}
)

func NewModelServerCaller() ModelServiceCaller {
	ModelServerCallerInit.Do(func() {
		modelServerCallerInstance = &modelServiceCaller{}
	})
	return modelServerCallerInstance
}

func (m modelServiceCaller) CallEmbedQuery(request *RequestEmbedQuery) ResponseEmbedQuery {
	req := jsonMarshall(request)
	res, err := http.Post(embedQueryUrl, "application/json", bytes.NewBuffer(req))
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	var responseDTO ResponseEmbedQuery
	parsingDto(&responseDTO, body)
	return responseDTO
}

func (m modelServiceCaller) CallQueryByIntent(query *RequestScriptQuery) ResponseScriptQuery {
	req := jsonMarshall(query)
	res, err := http.Post(getQueryByIntentUrl, "application/json", bytes.NewBuffer(req))
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	var responseDTO ResponseScriptQuery
	err = json.Unmarshal(body, &responseDTO)
	if err != nil {
		panic(err)
	}
	return responseDTO
}

func parsingDto(dto any, res []byte) {
	err := json.Unmarshal(res, dto)
	if err != nil {
		panic(err)
	}
}

func jsonMarshall(dto any) []byte {
	jsonData, err := json.Marshal(dto)
	if err != nil {
		panic(err)
	}
	return jsonData
}
