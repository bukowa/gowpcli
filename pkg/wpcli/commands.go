package wpcli

import (
	"github.com/bukowa/gowpcli/pkg/wpcli/core"
	"github.com/bukowa/gowpcli/pkg/wpcli/plugin"
)

type CoreInterface interface {
	Install(core.Install) Command
}

type PluginInterface interface {
	Activate(activate plugin.Activate) Command
}

type Core struct{}
type Plugin struct{}

func (p Plugin) Activate(activate plugin.Activate) Command {
	return &activate
}

func (c Core) Install(install core.Install) Command {
	return &install
}
