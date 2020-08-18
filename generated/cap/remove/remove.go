/*
## OPTIONS
	<role>
	: Key for the role.
	<cap>...
	: One or more capabilities to remove.
## EXAMPLES
	    # Remove 'spectate' capability from 'author' role.
	    $ wp cap remove author spectate
	    Success: Removed 1 capability from 'author' role.
	
 */
package remove
import utils "github.com/bukowa/gowpcli"

// Remove //Removes capabilities from a given role.
type Remove struct {
    Role string // <role>
    Cap []string // <cap>...
}

func (r Remove) Args() []string {
    var args = []string{"cap", "remove"}
    args = utils.MakeArg(args, "<role>", r.Role)
    args = utils.MakeArg(args, "<cap>...", r.Cap)
    return args
}

