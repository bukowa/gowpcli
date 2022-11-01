/*
## OPTIONS
	[<plugin>...]
	: One or more plugins to uninstall.
	[--deactivate]
	: Deactivate the plugin before uninstalling. Default behavior is to warn and skip if the plugin is active.
	[--skip-delete]
	: If set, the plugin files will not be deleted. Only the uninstall procedure
	will be run.
	[--all]
	: If set, all plugins will be uninstalled.
	[--exclude=<name>]
	: Comma separated list of plugin slugs to be excluded from uninstall.
## EXAMPLES
	    $ wp plugin uninstall hello
	    Uninstalled and deleted 'hello' plugin.
	    Success: Uninstalled 1 of 1 plugins.
	    # Uninstall all plugins excluding specified ones
	    $ wp plugin uninstall --all --exclude=hello-dolly,jetpack
	    Uninstalled and deleted 'akismet' plugin.
	    Uninstalled and deleted 'tinymce-templates' plugin.
	    Success: Uninstalled 2 of 2 plugins.
	
 */
package uninstall
import utils "github.com/bukowa/gowpcli"

// Uninstall //Uninstalls one or more plugins.
type Uninstall struct {
    Plugin []string // [<plugin>...]
    Deactivate bool // [--deactivate]
    SkipDelete bool // [--skip-delete]
    All bool // [--all]
    Exclude string // [--exclude=<name>]
}

func (u Uninstall) Args() []string {
    var args = []string{"plugin", "uninstall"}
    args = utils.MakeArg(args, "[<plugin>...]", u.Plugin)
    args = utils.MakeArg(args, "[--deactivate]", u.Deactivate)
    args = utils.MakeArg(args, "[--skip-delete]", u.SkipDelete)
    args = utils.MakeArg(args, "[--all]", u.All)
    args = utils.MakeArg(args, "[--exclude=<name>]", u.Exclude)
    return args
}

