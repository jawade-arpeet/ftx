package server

import (
	"fmt"
	"ftx/internal/client"
	"ftx/internal/config"
	"ftx/internal/handler"
	"ftx/internal/middleware"
	"ftx/internal/repository"
	"ftx/internal/router"
	"ftx/internal/service"

	"github.com/gin-gonic/gin"
)

type Server struct {
	config *config.ServerConfig
	router *gin.Engine
}

func New(
	cfg *config.ServerConfig,
	client *client.Client,
) *Server {
	repo := repository.New(client)
	service := service.New(repo)
	middleware := middleware.New()
	handler := handler.New(service)
	router := router.New(cfg.RunEnv, middleware, handler)

	return &Server{
		config: cfg,
		router: router,
	}
}

func (s *Server) Run() error {
	addr := fmt.Sprintf(":%d", s.config.Port)
	return s.router.Run(addr)
}
