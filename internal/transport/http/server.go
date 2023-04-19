package http

import (
	"awesomeProject/internal/config"
	"awesomeProject/internal/transport/http/handlers"
	"awesomeProject/internal/transport/http/middleware"
	"context"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type Server struct {
	cfg     *config.Config
	handler *handlers.Manager
	HTTP    *gin.Engine
}

func NewServer(cfg *config.Config, handler *handlers.Manager) *Server {
	return &Server{cfg: cfg, handler: handler}
}

func (s *Server) StartHTTpServer(ctx context.Context) error {
	r := gin.Default()
	r.Use(gin.Logger())
	s.UserRoutes(r)
	r.Use(middleware.Authentication())
	go func() {
		// service connections
		if err := r.Run(s.cfg.Port); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()
	<-ctx.Done()
	return nil
}
