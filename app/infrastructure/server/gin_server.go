package server

import (
	"context"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/pepese/golang-gin-sample/app"
	"github.com/pepese/golang-gin-sample/app/interface/controller"
	uuid "github.com/satori/go.uuid"
)

var ginServer *gin.Engine

func GinServerRun() {
	RunWithGracefulStop(ginServer)
}

func NewGinServer() *gin.Engine {
	e := gin.New()
	e.Use(accessLogging())
	e.Use(gin.Recovery())
	controller.NewGinRouter(e)
	ginServer = e

	return e
}

func accessLogging() gin.HandlerFunc {
	// Logging Skip Logic
	notlogged := []string{"/healthcheck"}
	var skip map[string]struct{}
	if length := len(notlogged); length > 0 {
		skip = make(map[string]struct{}, length)
		for _, path := range notlogged {
			skip[path] = struct{}{}
		}
	}
	return func(c *gin.Context) {
		// Request ID
		reqId := c.Request.Header.Get("X-Request-Id")
		if reqId == "" {
			reqId = c.Request.Header.Get("x-amzn-trace-id")
		}
		if reqId == "" {
			reqId = uuid.NewV4().String()
		}
		c.Set("reqId", reqId)
		c.Writer.Header().Set("X-Request-Id", reqId)

		// context.Context
		ctx := context.Background()

		// Logger
		logger := app.GetLoggerWithKeyValue("reqId", reqId)
		ctx = app.SetLoggerToContext(ctx, logger)
		c.Set("ctx", ctx)

		// Access Log
		start := time.Now()
		method := c.Request.Method
		path := c.Request.URL.Path
		query := c.Request.URL.RawQuery
		if query != "" {
			path = path + "?" + query
		}
		header := c.Request.Header

		c.Next()

		if _, ok := skip[path]; !ok {
			responseTime := time.Now().Sub(start)
			statusCode := c.Writer.Status()
			logger.Infow("access log", "method", method, "path", path, "header", fmt.Sprintf("%#v", header), "responseTime", responseTime, "statusCode", statusCode)
		}
	}
}
