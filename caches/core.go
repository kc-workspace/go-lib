package caches

import (
	"sync"

	"github.com/kc-workspace/go-lib/logger"
)

func New() *Service {
	return &Service{
		caches: make(map[string]*Data),
		mutex:  sync.RWMutex{},
		logger: logger.Get("cache", "service"),
	}
}

var Global = New()
