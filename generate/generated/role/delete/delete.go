/*
## OPTIONS
	<role-key>
	: The internal name of the role.
## EXAMPLES
	    # Delete approver role.
	    $ wp role delete approver
	    Success: Role with key 'approver' deleted.
	    # Delete productadmin role.
	    wp role delete productadmin
	    Success: Role with key 'productadmin' deleted.
	
 */
package delete
import utils "github.com/bukowa/gowpcli"

// Delete //Deletes an existing role.
type Delete struct {
    RoleKey string // <role-key>
}

func (d Delete) Args() []string {
    var args = []string{"role", "delete"}
    args = utils.MakeArg(args, "<role-key>", d.RoleKey)
    return args
}

