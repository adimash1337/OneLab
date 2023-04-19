package handlers

import (
	"awesomeProject/internal/config"
	"awesomeProject/internal/service"
)

type Manager struct {
	srv *service.Manager
}

func NewManager(cfg *config.Config, srv *service.Manager) *Manager {
	return &Manager{srv: srv}
}
