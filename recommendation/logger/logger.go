package logger

import (
	"bytes"
	"io"
	"os"
	"time"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
)

var logger *lumberjack.Logger

func SetupLogger(filename string, maxsize int, maxbackup int, compress bool, level string, client *elasticsearch.Client) {
	lvl, err := log.ParseLevel(level)
	if err != nil {
		log.SetLevel(lvl)
	} else {
		log.SetLevel(log.InfoLevel)
	}

	// use lumberjack to write to implement rotation.
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.InfoLevel)

	// use elasticsearch hooking
	log.AddHook(NewElasticsearchHook(client, "log"))

	// set up request logging
}

func Close() {
	logger.Close()
}

func WithTrace(ctx *gin.Context) *log.Entry {
	fields := log.Fields{}
	if traceID := ctx.Request.Header.Get("X-Trace-ID"); traceID != "" {
		fields["trace_id"] = traceID
	}
	if spanID := ctx.Request.Header.Get("X-Span-ID"); spanID != "" {
		fields["span_id"] = spanID
	}

	return log.WithFields(fields)
}

func SetLoggerHooking(r *gin.Engine) {
	r.Use(func(c *gin.Context) {
		start := time.Now()
		c.Next()
		latency := time.Since(start)
		status := c.Writer.Status()

		// Log request details
		fields := log.Fields{
			"method":  c.Request.Method,
			"path":    c.Request.URL.Path,
			"status":  status,
			"latency": latency,
		}

		// Add trace and span IDs if available
		if traceEntry := WithTrace(c); traceEntry != nil {
			for k, v := range traceEntry.Data {
				traceEntry = traceEntry.WithField(k, v)
			}
		}
		log.WithFields(fields).Info("request completed")
	})
}

func RequestLoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		// Clone the request body for logging and processing
		var requestBodyBytes []byte
		if c.Request.Body != nil {
			// Read and save the request body
			requestBodyBytes, _ = io.ReadAll(c.Request.Body)

			// Restore the original request body for downstream handlers
			c.Request.Body = io.NopCloser(bytes.NewBuffer(requestBodyBytes))
		}

		// Process the request
		c.Next()

		// Log request details
		log.WithFields(log.Fields{
			"method":      c.Request.Method,
			"path":        c.Request.URL.Path,
			"status":      c.Writer.Status(),
			"latency":     time.Since(start).String(),
			"queryParams": c.Request.URL.Query(),
			"headers":     c.Request.Header,
			"body":        string(requestBodyBytes),
		}).Info("API Request")
	}
}
