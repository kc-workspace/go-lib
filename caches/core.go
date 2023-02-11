package caches

import (
	"sync"

	"github.com/kc-workspace/go-lib/logger"
)

func New(l *logger.Logger) *Service {
	return &Service{
		caches: make(map[string]*Data),
		mutex:  sync.RWMutex{},
		logger: l,
	}
}

var Global = New(logger.DefaultManager.New("caches", "global"))
