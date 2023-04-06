package http

import (
	"awesomeProject/internal/config"
	"awesomeProject/internal/logger"
	"awesomeProject/internal/transport/http/handlers"
	"context"
	"github.com/labstack/echo/v4"
	"net/http"
	"os"
	"os/signal"
	"time"
)

type Server struct {
	cfg     *config.Config
	handler *handlers.UserHandler
	HTTP    *echo.Echo
}

func NewServer(cfg *config.Config, handler *handlers.UserHandler) *Server {
	return &Server{cfg: cfg, handler: handler}
}

func (s *Server) StartHTTpServer(ctx context.Context) error {
	s.HTTP = echo.New()
	s.Router()
	if err := s.HTTP.Start(s.cfg.Port); err != http.ErrServerClosed {
		logger.Logger().Println(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	gracefulShutdown(cancel)

	if err := s.HTTP.Shutdown(ctx); err != nil {
		logger.Logger().Println(err)
		s.HTTP.Logger.Fatal(err)
	}
	return nil
}

func gracefulShutdown(c context.CancelFunc) {
	osC := make(chan os.Signal, 1)
	signal.Notify(osC, os.Interrupt)
	go func() {
		logger.Logger().Println(<-osC)
	}()
}
