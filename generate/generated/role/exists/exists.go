/*
## INFO
	Exits with return code 0 if the role exists, 1 if it does not.
## OPTIONS
	<role-key>
	: The internal name of the role.
## EXAMPLES
	    # Check if a role exists.
	    $ wp role exists editor
	    Success: Role with ID 'editor' exists.
	
 */
package exists
import utils "github.com/bukowa/gowpcli"

// Exists //Checks if a role exists.
type Exists struct {
    RoleKey string // <role-key>
}

func (e Exists) Args() []string {
    var args = []string{"role", "exists"}
    args = utils.MakeArg(args, "<role-key>", e.RoleKey)
    return args
}

