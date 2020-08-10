package plugin

import (
	"github.com/bukowa/gowpcli/pkg/wpcli/utils"
)

type Activate struct {
	Plugin  string
	All     bool
	Network bool
}

func (cmd *Activate) Args() (args []string) {
	var command = []string{"plugin", "activate"}

	args = append(args, command...)
	if cmd.Plugin != "" {
		args = utils.MakeArg(args, "plugin", cmd.Plugin)
	}
	if cmd.All {
		args = utils.MakeArg(args, "all", "")
	}
	if cmd.Network {
		args = utils.MakeArg(args, "network", "")
	}
	return args
}
