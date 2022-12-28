package plugins

import (
	"github.com/kc-workspace/go-lib/commandline/hooks"
	"github.com/kc-workspace/go-lib/logger"
	"github.com/kc-workspace/go-lib/mapper"
)

// SupportBanner will create application info banner
func SupportBanner(p *PluginParameter) error {
	p.NewHook(hooks.BEFORE_COMMAND, func(m mapper.Mapper) error {
		if logger.GetLevel() == logger.DEBUG {
			p.Logger.Debug("%-12s: %s", "metadata", p.Metadata.String())
			p.Logger.Debug("%-12s: %v", "config", m.String())

			return nil
		} else {
			var table = logger.GetTable(2)
			table.Row("Name", "Version", "Commit")
			table.Row(p.Metadata.Name, p.Metadata.Version, p.Metadata.Commit)
			return table.End()
		}
	})
	return nil
}
