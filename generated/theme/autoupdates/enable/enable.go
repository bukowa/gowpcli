/*
## OPTIONS
	[<theme>...]
	: One or more themes to enable auto-updates for.
	[--all]
	: If set, auto-updates will be enabled for all themes.
	[--disabled-only]
	: If set, filters list of themes to only include the ones that have
	auto-updates disabled.
## EXAMPLES
	    # Enable the auto-updates for a theme
	    $ wp theme auto-updates enable twentysixteen
	    Theme auto-updates for 'twentysixteen' enabled.
	    Success: Enabled 1 of 1 theme auto-updates.
	
 */
package enable
import utils "github.com/bukowa/gowpcli"

// Enable //Enables the auto-updates for a theme.
type Enable struct {
    Theme []string // [<theme>...]
    All bool // [--all]
    DisabledOnly bool // [--disabled-only]
}

func (e Enable) Args() []string {
    var args = []string{"theme", "auto-updates", "enable"}
    args = utils.MakeArg(args, "[<theme>...]", e.Theme)
    args = utils.MakeArg(args, "[--all]", e.All)
    args = utils.MakeArg(args, "[--disabled-only]", e.DisabledOnly)
    return args
}

