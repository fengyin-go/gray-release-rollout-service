// Package app 负责依赖装配。
package app

import (
	"net/http"

	"grayrelease/internal/config"
	"grayrelease/internal/handler"
	"grayrelease/internal/service"
	"grayrelease/internal/store"
	"grayrelease/pkg/logger"
)

// App 组装完成的应用。
type App struct {
	server *handler.Server
}

// New 装配 store -> service -> handler。
func New(cfg *config.Config, log *logger.Logger) (*App, error) {
	st := store.NewMemoryStore()
	svc := service.New(st, log, cfg)
	server := handler.NewServer(svc, log, cfg)
	log.Infof("应用装配完成，配置：%s", cfg.String())
	return &App{server: server}, nil
}

// Routes 返回 HTTP 路由。
func (a *App) Routes() http.Handler { return a.server.Routes() }
