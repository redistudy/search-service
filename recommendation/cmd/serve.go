package cmd

import (
	"context"
	"net"
	"net/http"
	"path/filepath"
	"recommendation/logger"
	"recommendation/server"
	"recommendation/setting"
	"recommendation/signals"
	"time"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "search-demo-template",
	Short: "A clean architecture template for Golang Gin services",

	Run: func(cmd *cobra.Command, args []string) {
		var config setting.Configuration

		if err := setting.LoadConfig(&config); err != nil {
			log.Error("loading config file failed.", err)
			return
		}

		gin.SetMode(config.Server.RunMode)

		// init db
		engine, err := dbInit(&config.Database)
		if err != nil {
			log.Error("init db failed.", err)
			return
		}

		// init recommendation db
		recommendationEngine, err := recommendationDbInit(&config.Database)
		if err != nil {
			log.Error("init recommendation db failed.", err)
			return
		}

		// init logger
		logger.SetupLogger(
			filepath.Join(
				config.Log.LogSavePath,
				config.Log.LogFileName),
			config.Log.MaxSize,
			config.Log.MaxBackups,
			config.Log.Compress,
			config.Log.Level, engine)

		// start http server
		svr := server.NewServer(&config, engine, recommendationEngine)
		if err := svr.Start(); err != nil {
			log.Error("init server failed.", err)
			return
		}

		// graceful shutdown
		stopCh := signals.SetupSignalHandler()
		sd, _ := signals.NewShutdown(config.App.ServerShutdownTimeout)
		sd.Graceful(stopCh, svr, engine)
	},
}

func Execute() error {
	rootCmd.CompletionOptions.DisableDefaultCmd = true

	return rootCmd.Execute()
}

func dbInit(dbc *setting.DatabaseSettingS) (*elasticsearch.Client, error) {
	// var err error
	config := elasticsearch.Config{
		Addresses: []string{
			dbc.Address,
		},
		Transport: &http.Transport{
			MaxIdleConnsPerHost:   dbc.IdleHost,
			ResponseHeaderTimeout: time.Second,
			DialContext:           (&net.Dialer{Timeout: time.Second}).DialContext,
		},
	}
	return elasticsearch.NewClient(config)
}

func recommendationDbInit(dbc *setting.DatabaseSettingS) (*redis.Client, error) {
	// var err error
	client := redis.NewClient(&redis.Options{
		Addr:     dbc.Redis.Address,
		DB:       dbc.Redis.Db,
		Username: dbc.Redis.Username,
		Password: dbc.Redis.Password,
	})
	if err := client.Ping(context.Background()).Err(); err != nil {
		return nil, err
	}
	return client, nil
}
