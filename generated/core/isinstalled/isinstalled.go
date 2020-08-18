/*
## INFO
	Determines whether WordPress is installed by checking if the standard
	database tables are installed. Doesn't produce output; uses exit codes
	to communicate whether WordPress is installed.
	[--network]
	: Check if this is a multisite installation.
## EXAMPLES
	    # Check whether WordPress is installed; exit status 0 if installed, otherwise 1
	    $ wp core is-installed
	    $ echo $?
	    1
	    # Bash script for checking whether WordPress is installed or not
	    if ! $(wp core is-installed); then
	        wp core install
	    fi
	
 */
package isinstalled
import utils "github.com/bukowa/gowpcli"

// IsInstalled //Checks if WordPress is installed.
type IsInstalled struct {
    Network bool // [--network]
}

func (i IsInstalled) Args() []string {
    var args = []string{"core", "is-installed"}
    args = utils.MakeArg(args, "[--network]", i.Network)
    return args
}

