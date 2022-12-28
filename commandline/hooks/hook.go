package hooks

import "github.com/kc-workspace/go-lib/mapper"

// Hook action
type Hook func(config mapper.Mapper) error
