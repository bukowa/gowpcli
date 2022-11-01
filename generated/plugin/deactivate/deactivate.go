/*
## OPTIONS
	[<plugin>...]
	: One or more plugins to deactivate.
	[--uninstall]
	: Uninstall the plugin after deactivation.
	[--all]
	: If set, all plugins will be deactivated.
	 [--exclude=<name>]
	: Comma separated list of plugin slugs that should be excluded from deactivation.
	[--network]
	: If set, the plugin will be deactivated for the entire multisite network.
## EXAMPLES
	    # Deactivate plugin
	    $ wp plugin deactivate hello
	    Plugin 'hello' deactivated.
	    Success: Deactivated 1 of 1 plugins.
	    # Deactivate all plugins with exclusion
	    $ wp plugin deactivate --all --exclude=hello,wordpress-seo
	    Plugin 'contact-form-7' deactivated.
	    Plugin 'ninja-forms' deactivated.
	    Success: Deactivated 2 of 2 plugins.
	
 */
package deactivate
import utils "github.com/bukowa/gowpcli"

// Deactivate //Deactivates one or more plugins.
type Deactivate struct {
    Plugin []string // [<plugin>...]
    Uninstall bool // [--uninstall]
    All bool // [--all]
    Exclude string // [--exclude=<name>]
    Network bool // [--network]
}

func (d Deactivate) Args() []string {
    var args = []string{"plugin", "deactivate"}
    args = utils.MakeArg(args, "[<plugin>...]", d.Plugin)
    args = utils.MakeArg(args, "[--uninstall]", d.Uninstall)
    args = utils.MakeArg(args, "[--all]", d.All)
    args = utils.MakeArg(args, "[--exclude=<name>]", d.Exclude)
    args = utils.MakeArg(args, "[--network]", d.Network)
    return args
}

