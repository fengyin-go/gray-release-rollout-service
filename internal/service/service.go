package service

import (
	"grayrelease/internal/config"
	"grayrelease/internal/store"
	"grayrelease/pkg/logger"
)

// Service 聚合业务逻辑，依赖 Store 与配置。
type Service struct {
	store store.Store
	log   *logger.Logger
	cfg   *config.Config
}

// New 构造服务实例。
func New(st store.Store, log *logger.Logger, cfg *config.Config) *Service {
	return &Service{store: st, log: log, cfg: cfg}
}
