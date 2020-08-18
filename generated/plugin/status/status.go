/*
## OPTIONS
	[<plugin>]
	: A particular plugin to show the status for.
## EXAMPLES
	    # Displays status of all plugins
	    $ wp plugin status
	    5 installed plugins:
	      I akismet                3.1.11
	      I easy-digital-downloads 2.5.16
	      A theme-check            20160523.1
	      I wen-logo-slider        2.0.3
	      M ns-pack                1.0.0
	    Legend: I = Inactive, A = Active, M = Must Use
	    # Displays status of a plugin
	    $ wp plugin status theme-check
	    Plugin theme-check details:
	        Name: Theme Check
	        Status: Active
	        Version: 20160523.1
	        Author: Otto42, pross
	        Description: A simple and easy way to test your theme for all the latest WordPress standards and practices. A great theme development tool!
	
 */
package status
import utils "github.com/bukowa/gowpcli"

// Status //Reveals the status of one or all plugins.
type Status struct {
    Plugin string // [<plugin>]
}

func (s Status) Args() []string {
    var args = []string{"plugin", "status"}
    args = utils.MakeArg(args, "[<plugin>]", s.Plugin)
    return args
}

