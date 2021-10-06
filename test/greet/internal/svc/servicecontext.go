package svc

import (
	"github.com/muhfajar/go-zero-cors-middleware/test/greet/internal/config"
)

type ServiceContext struct {
	Config config.Config
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
	}
}
