/*
## INFO
	Returns exit code 0 when active, 1 when not active.
## OPTIONS
	<theme>
	: The plugin to check.
## EXAMPLES
	    # Check whether theme is Active; exit status 0 if active, otherwise 1
	    $ wp theme is-active twentyfifteen
	    $ echo $?
	    1
	
 */
package isactive
import utils "github.com/bukowa/gowpcli"

// IsActive //Checks if a given theme is active.
type IsActive struct {
    Theme string // <theme>
}

func (i IsActive) Args() []string {
    var args = []string{"theme", "is-active"}
    args = utils.MakeArg(args, "<theme>", i.Theme)
    return args
}

