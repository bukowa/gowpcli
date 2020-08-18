/*
## OPTIONS
	<role>
	: Key for the role.
	<cap>...
	: One or more capabilities to add.
	[--grant]
	: Adds the capability as an explicit boolean value, instead of implicitly defaulting to `true`.
	---
	default: true
	options:
	  - true
	  - false
	---
## EXAMPLES
	    # Add 'spectate' capability to 'author' role.
	    $ wp cap add author spectate
	    Success: Added 1 capability to 'author' role.
	
 */
package add
import utils "github.com/bukowa/gowpcli"

// Add //Adds capabilities to a given role.
type Add struct {
    Role string // <role>
    Cap []string // <cap>...
    Grant bool // [--grant]
}

func (a Add) Args() []string {
    var args = []string{"cap", "add"}
    args = utils.MakeArg(args, "<role>", a.Role)
    args = utils.MakeArg(args, "<cap>...", a.Cap)
    args = utils.MakeArg(args, "[--grant]", a.Grant)
    return args
}

