# Recommendation Service

## Overview
Go-based search service using the Gin framework for building RESTful APIs.
[intent or embedding using fastapi](/fastapi/intent-recognition-api/README.md)

## Tech Stack
- Go 1.19+
- Gin Web Framework
- Elasticsearch
- Redis Cache

## Project Structure
```
/recommendation
├── README.md
├── api
│   ├── router.go
│   └── v1
│       └── poi.go
├── app
│   ├── app.go
│   └── recommendation
├── cmd
│   └── serve.go
├── configs
│   └── config.yaml
├── docker-compose.yml
├── domain
│   └── response.go
├── dto
│   └── poiDto.go
├── elasticsearch
│   └── docker-compose.yml
├── go.mod
├── go.sum
├── infrastructure
│   ├── ModelServerCaller.go
│   ├── dto.go
│   └── infratructrue.go
├── internal
│   ├── elasticsearch-client
│   │   ├── PoiDto.go
│   │   ├── elasticsearchClient.go
│   │   └── latlonConverter.go
│   ├── redis-client
│   │   └── RedisClient.go
│   └── service
│       └── poiSaveService.go
├── logger
│   ├── logger.go
│   └── loggerHooker.go
├── main.go
├── recommendation
├── repository
│   ├── http-client.private.env.json
│   └── poiRepository.go
├── server
│   └── server.go
├── service
│   ├── Poi.go
│   └── searchService.go
├── setting
│   └── setting.go
└── signals
    └── shutdown.go

```

## Getting Started
1. Clone the repository:
    
2. Navigate to the project directory:
    ```sh
    cd recommendation-service
    ```
3. Install dependencies:
    ```sh
    go mod tidy
    ```
4. Run the application:
    ```sh
    go run cmd/main.go
    ```

## Endpoints
- `POST /poi` - search by text

## License
This project is licensed under the MIT License.