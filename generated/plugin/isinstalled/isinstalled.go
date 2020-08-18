/*
## INFO
	Returns exit code 0 when installed, 1 when uninstalled.
## OPTIONS
	<plugin>
	: The plugin to check.
## EXAMPLES
	    # Check whether plugin is installed; exit status 0 if installed, otherwise 1
	    $ wp plugin is-installed hello
	    $ echo $?
	    1
	
 */
package isinstalled
import utils "github.com/bukowa/gowpcli"

// IsInstalled //Checks if a given plugin is installed.
type IsInstalled struct {
    Plugin string // <plugin>
}

func (i IsInstalled) Args() []string {
    var args = []string{"plugin", "is-installed"}
    args = utils.MakeArg(args, "<plugin>", i.Plugin)
    return args
}

