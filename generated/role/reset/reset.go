/*
## OPTIONS
	[<role-key>...]
	: The internal name of one or more roles to reset.
	[--all]
	: If set, all default roles will be reset.
## EXAMPLES
	    # Reset role.
	    $ wp role reset administrator author contributor
	    Success: Reset 1/3 roles.
	    # Reset all default roles.
	    $ wp role reset --all
	    Success: All default roles reset.
	
 */
package reset
import utils "github.com/bukowa/gowpcli"

// Reset //Resets any default role to default capabilities.
type Reset struct {
    RoleKey []string // [<role-key>...]
    All bool // [--all]
}

func (r Reset) Args() []string {
    var args = []string{"role", "reset"}
    args = utils.MakeArg(args, "[<role-key>...]", r.RoleKey)
    args = utils.MakeArg(args, "[--all]", r.All)
    return args
}

