/*
## INFO
	Permits theme to be activated from the dashboard of a site on a WordPress
	multisite install.
## OPTIONS
	<theme>
	: The theme to enable.
	[--network]
	: If set, the theme is enabled for the entire network
	[--activate]
	: If set, the theme is activated for the current site. Note that
	the "network" flag has no influence on this.
## EXAMPLES
	    # Enable theme
	    $ wp theme enable twentysixteen
	    Success: Enabled the 'Twenty Sixteen' theme.
	    # Network enable theme
	    $ wp theme enable twentysixteen --network
	    Success: Network enabled the 'Twenty Sixteen' theme.
	    # Network enable and activate theme for current site
	    $ wp theme enable twentysixteen --activate
	    Success: Enabled the 'Twenty Sixteen' theme.
	    Success: Switched to 'Twenty Sixteen' theme.
	
 */
package enable
import utils "github.com/bukowa/gowpcli"

// Enable //Enables a theme on a WordPress multisite install.
type Enable struct {
    Theme string // <theme>
    Network bool // [--network]
    Activate bool // [--activate]
}

func (e Enable) Args() []string {
    var args = []string{"theme", "enable"}
    args = utils.MakeArg(args, "<theme>", e.Theme)
    args = utils.MakeArg(args, "[--network]", e.Network)
    args = utils.MakeArg(args, "[--activate]", e.Activate)
    return args
}

