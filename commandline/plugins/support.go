package plugins

// Support is example how we can implement plugins support
func Support() Plugin {
	return func(p *PluginParameter) error {
		return nil
	}
}
