package server

import (
	"business-auth/conf/http_server"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

const defaultHost = "0.0.0.0"

type HttpServer interface {
	Start()
	Stop()
}

type httpServer struct {
	Port   uint
	server *http.Server
}

func NewHttpServer(router *gin.Engine, config http_server.HttpServerConfig) HttpServer {
	return &httpServer{
		Port: config.Port,
		server: &http.Server{
			Addr:    fmt.Sprintf("%s:%d", defaultHost, config.Port),
			Handler: router,
		},
	}
}

func (httpServer httpServer) Start() {
	go func() {
		if err := httpServer.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf(
				"failed to stater HttpServer listen port %d, err=%s",
				httpServer.Port, err.Error(),
			)
		}
	}()
	log.Printf("Start Service with port %d", httpServer.Port)
}

func (httpServer httpServer) Stop() {
	ctx, cancel := context.WithTimeout(
		context.Background(), time.Duration(3)*time.Second,
	)
	defer cancel()

	if err := httpServer.server.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown err=%s", err.Error())
	}
}
