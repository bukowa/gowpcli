/*
## OPTIONS
	[--extra]
	: Show extended version information.
## EXAMPLES
	    # Display the WordPress version
	    $ wp core version
	    4.5.2
	    # Display WordPress version along with other information
	    $ wp core version --extra
	    WordPress version: 4.5.2
	    Database revision: 36686
	    TinyMCE version:   4.310 (4310-20160418)
	    Package language:  en_US
	
 */
package version
import utils "github.com/bukowa/gowpcli"

// Version //Displays the WordPress version.
type Version struct {
    Extra bool // [--extra]
}

func (v Version) Args() []string {
    var args = []string{"core", "version"}
    args = utils.MakeArg(args, "[--extra]", v.Extra)
    return args
}

