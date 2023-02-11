package plugins

import (
	"github.com/kc-workspace/go-lib/commandline/hooks"
	"github.com/kc-workspace/go-lib/mapper"
)

// SupportBanner will create application info banner
func SupportBanner() Plugin {
	return func(p *PluginParameter) error {
		p.NewHook(hooks.BEFORE_COMMAND, func(m mapper.Mapper) error {
			if p.Logger.IsDebug() {
				p.Logger.Debug("%-12s: %s", "metadata", p.Metadata.String())
				p.Logger.Debug("%-12s: %v", "config", m.String())

				return nil
			} else {
				var table = p.Logger.ToTable(3)
				return table.
					Row("Name", "Version", "Commit").
					Row(p.Metadata.Name, p.Metadata.Version, p.Metadata.Commit).
					End()
			}
		})
		return nil
	}
}
