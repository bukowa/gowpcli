/*
## OPTIONS
	<role-key>
	: The internal name of the role.
	<role-name>
	: The publicly visible name of the role.
	[--clone=<role>]
	: Clone capabilities from an existing role.
## EXAMPLES
	    # Create role for Approver.
	    $ wp role create approver Approver
	    Success: Role with key 'approver' created.
	    # Create role for Product Administrator.
	    $ wp role create productadmin "Product Administrator"
	    Success: Role with key 'productadmin' created.
	
 */
package create
import utils "github.com/bukowa/gowpcli"

// Create //Creates a new role.
type Create struct {
    RoleKey string // <role-key>
    RoleName string // <role-name>
    Clone string // [--clone=<role>]
}

func (c Create) Args() []string {
    var args = []string{"role", "create"}
    args = utils.MakeArg(args, "<role-key>", c.RoleKey)
    args = utils.MakeArg(args, "<role-name>", c.RoleName)
    args = utils.MakeArg(args, "[--clone=<role>]", c.Clone)
    return args
}

