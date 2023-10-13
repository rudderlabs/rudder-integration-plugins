package sources

import (
	"github.com/rudderlabs/rudder-plugins-manager/plugins"
)

func NewSourceManager() *plugins.BasePluginManager {
	manager := plugins.NewBasePluginManager()
	manager.Add(webHookPlugin)
	return manager
}
