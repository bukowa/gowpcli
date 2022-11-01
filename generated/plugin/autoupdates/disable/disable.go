/*
## OPTIONS
	[<plugin>...]
	: One or more plugins to disable auto-updates for.
	[--all]
	: If set, auto-updates will be disabled for all plugins.
	[--enabled-only]
	: If set, filters list of plugins to only include the ones that have
	auto-updates enabled.
## EXAMPLES
	    # Disable the auto-updates for a plugin
	    $ wp plugin auto-updates disable hello
	    Plugin auto-updates for 'hello' disabled.
	    Success: Disabled 1 of 1 plugin auto-updates.
	
 */
package disable
import utils "github.com/bukowa/gowpcli"

// Disable //Disables the auto-updates for a plugin.
type Disable struct {
    Plugin []string // [<plugin>...]
    All bool // [--all]
    EnabledOnly bool // [--enabled-only]
}

func (d Disable) Args() []string {
    var args = []string{"plugin", "auto-updates", "disable"}
    args = utils.MakeArg(args, "[<plugin>...]", d.Plugin)
    args = utils.MakeArg(args, "[--all]", d.All)
    args = utils.MakeArg(args, "[--enabled-only]", d.EnabledOnly)
    return args
}

