/*
## OPTIONS
	<user>
	: User ID, user email, or user login.
	[<role>]
	: A specific role to remove.
## EXAMPLES
	    $ wp user remove-role 12 author
	    Success: Removed 'author' role for johndoe (12).
	
 */
package removerole
import utils "github.com/bukowa/gowpcli"

// RemoveRole //Removes a user's role.
type RemoveRole struct {
    User string // <user>
    Role string // [<role>]
}

func (r RemoveRole) Args() []string {
    var args = []string{"user", "remove-role"}
    args = utils.MakeArg(args, "<user>", r.User)
    args = utils.MakeArg(args, "[<role>]", r.Role)
    return args
}

