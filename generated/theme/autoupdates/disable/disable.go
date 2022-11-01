/*
## OPTIONS
	[<theme>...]
	: One or more themes to disable auto-updates for.
	[--all]
	: If set, auto-updates will be disabled for all themes.
	[--enabled-only]
	: If set, filters list of themes to only include the ones that have
	auto-updates enabled.
## EXAMPLES
	    # Disable the auto-updates for a theme
	    $ wp theme auto-updates disable twentysixteen
	    Theme auto-updates for 'twentysixteen' disabled.
	    Success: Disabled 1 of 1 theme auto-updates.
	
 */
package disable
import utils "github.com/bukowa/gowpcli"

// Disable //Disables the auto-updates for a theme.
type Disable struct {
    Theme []string // [<theme>...]
    All bool // [--all]
    EnabledOnly bool // [--enabled-only]
}

func (d Disable) Args() []string {
    var args = []string{"theme", "auto-updates", "disable"}
    args = utils.MakeArg(args, "[<theme>...]", d.Theme)
    args = utils.MakeArg(args, "[--all]", d.All)
    args = utils.MakeArg(args, "[--enabled-only]", d.EnabledOnly)
    return args
}

