/*
## INFO
	If the plugin is active, then it will be deactivated. If the plugin is
	inactive, then it will be activated.
## OPTIONS
	<plugin>...
	: One or more plugins to toggle.
	[--network]
	: If set, the plugin will be toggled for the entire multisite network.
## EXAMPLES
	    # Akismet is currently activated
	    $ wp plugin toggle akismet
	    Plugin 'akismet' deactivated.
	    Success: Toggled 1 of 1 plugins.
	    # Akismet is currently deactivated
	    $ wp plugin toggle akismet
	    Plugin 'akismet' activated.
	    Success: Toggled 1 of 1 plugins.
	
 */
package toggle
import utils "github.com/bukowa/gowpcli"

// Toggle //Toggles a plugin's activation state.
type Toggle struct {
    Plugin []string // <plugin>...
    Network bool // [--network]
}

func (t Toggle) Args() []string {
    var args = []string{"plugin", "toggle"}
    args = utils.MakeArg(args, "<plugin>...", t.Plugin)
    args = utils.MakeArg(args, "[--network]", t.Network)
    return args
}

