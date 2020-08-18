/*
## OPTIONS
	<user>
	: User ID, user email, or user login.
	<role>
	: Add the specified role to the user.
## EXAMPLES
	    $ wp user add-role 12 author
	    Success: Added 'author' role for johndoe (12).
	
 */
package addrole
import utils "github.com/bukowa/gowpcli"

// AddRole //Adds a role for a user.
type AddRole struct {
    User string // <user>
    Role string // <role>
}

func (a AddRole) Args() []string {
    var args = []string{"user", "add-role"}
    args = utils.MakeArg(args, "<user>", a.User)
    args = utils.MakeArg(args, "<role>", a.Role)
    return args
}

