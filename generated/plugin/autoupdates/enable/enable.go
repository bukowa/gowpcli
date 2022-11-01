/*
## OPTIONS
	[<plugin>...]
	: One or more plugins to enable auto-updates for.
	[--all]
	: If set, auto-updates will be enabled for all plugins.
	[--disabled-only]
	: If set, filters list of plugins to only include the ones that have
	auto-updates disabled.
## EXAMPLES
	    # Enable the auto-updates for a plugin
	    $ wp plugin auto-updates enable hello
	    Plugin auto-updates for 'hello' enabled.
	    Success: Enabled 1 of 1 plugin auto-updates.
	
 */
package enable
import utils "github.com/bukowa/gowpcli"

// Enable //Enables the auto-updates for a plugin.
type Enable struct {
    Plugin []string // [<plugin>...]
    All bool // [--all]
    DisabledOnly bool // [--disabled-only]
}

func (e Enable) Args() []string {
    var args = []string{"plugin", "auto-updates", "enable"}
    args = utils.MakeArg(args, "[<plugin>...]", e.Plugin)
    args = utils.MakeArg(args, "[--all]", e.All)
    args = utils.MakeArg(args, "[--disabled-only]", e.DisabledOnly)
    return args
}

