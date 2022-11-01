/*
## OPTIONS
	[<theme>...]
	: One or more themes to show the status of the auto-updates of.
	[--all]
	: If set, the status of auto-updates for all themes will be shown.
	[--enabled-only]
	: If set, filters list of themes to only include the ones that have
	auto-updates enabled.
	[--disabled-only]
	: If set, filters list of themes to only include the ones that have
	auto-updates disabled.
	[--field=<field>]
	: Only show the provided field.
## EXAMPLES
	    # Get the status of theme auto-updates
	    $ wp theme auto-updates status twentysixteen
	    +---------------+----------+
	    | name          | status   |
	    +---------------+----------+
	    | twentysixteen | disabled |
	    +---------------+----------+
	    # Get the list of themes that have auto-updates enabled
	    $ wp theme auto-updates status --all --enabled-only --field=name
	    twentysixteen
	    twentyseventeen
	
 */
package status
import utils "github.com/bukowa/gowpcli"

// Status //Shows the status of auto-updates for a theme.
type Status struct {
    Theme []string // [<theme>...]
    All bool // [--all]
    EnabledOnly bool // [--enabled-only]
    DisabledOnly bool // [--disabled-only]
    Field string // [--field=<field>]
}

func (s Status) Args() []string {
    var args = []string{"theme", "auto-updates", "status"}
    args = utils.MakeArg(args, "[<theme>...]", s.Theme)
    args = utils.MakeArg(args, "[--all]", s.All)
    args = utils.MakeArg(args, "[--enabled-only]", s.EnabledOnly)
    args = utils.MakeArg(args, "[--disabled-only]", s.DisabledOnly)
    args = utils.MakeArg(args, "[--field=<field>]", s.Field)
    return args
}

