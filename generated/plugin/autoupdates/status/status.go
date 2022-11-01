/*
## OPTIONS
	[<plugin>...]
	: One or more plugins to show the status of the auto-updates of.
	[--all]
	: If set, the status of auto-updates for all plugins will be shown.
	[--enabled-only]
	: If set, filters list of plugins to only include the ones that have
	auto-updates enabled.
	[--disabled-only]
	: If set, filters list of plugins to only include the ones that have
	auto-updates disabled.
	[--field=<field>]
	: Only show the provided field.
## EXAMPLES
	    # Get the status of plugin auto-updates
	    $ wp plugin auto-updates status hello
	    +-------+----------+
	    | name  | status   |
	    +-------+----------+
	    | hello | disabled |
	    +-------+----------+
	    # Get the list of plugins that have auto-updates enabled
	    $ wp plugin auto-updates status --all --enabled-only --field=name
	    akismet
	    duplicate-post
	
 */
package status
import utils "github.com/bukowa/gowpcli"

// Status //Shows the status of auto-updates for a plugin.
type Status struct {
    Plugin []string // [<plugin>...]
    All bool // [--all]
    EnabledOnly bool // [--enabled-only]
    DisabledOnly bool // [--disabled-only]
    Field string // [--field=<field>]
}

func (s Status) Args() []string {
    var args = []string{"plugin", "auto-updates", "status"}
    args = utils.MakeArg(args, "[<plugin>...]", s.Plugin)
    args = utils.MakeArg(args, "[--all]", s.All)
    args = utils.MakeArg(args, "[--enabled-only]", s.EnabledOnly)
    args = utils.MakeArg(args, "[--disabled-only]", s.DisabledOnly)
    args = utils.MakeArg(args, "[--field=<field>]", s.Field)
    return args
}

