/*
## INFO
	Returns exit code 0 when installed, 1 when uninstalled.
## OPTIONS
	<theme>
	: The theme to check.
## EXAMPLES
	    # Check whether theme is installed; exit status 0 if installed, otherwise 1
	    $ wp theme is-installed hello
	    $ echo $?
	    1
	
 */
package isinstalled
import utils "github.com/bukowa/gowpcli"

// IsInstalled //Checks if a given theme is installed.
type IsInstalled struct {
    Theme string // <theme>
}

func (i IsInstalled) Args() []string {
    var args = []string{"theme", "is-installed"}
    args = utils.MakeArg(args, "<theme>", i.Theme)
    return args
}

