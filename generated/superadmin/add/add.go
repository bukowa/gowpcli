/*
## OPTIONS
	<user>...
	: One or more user IDs, user emails, or user logins.
## EXAMPLES
	    $ wp super-admin add superadmin2
	    Success: Granted super-admin capabilities.
	
 */
package add
import utils "github.com/bukowa/gowpcli"

// Add //Grants super admin privileges to one or more users.
type Add struct {
    User []string // <user>...
}

func (a Add) Args() []string {
    var args = []string{"super-admin", "add"}
    args = utils.MakeArg(args, "<user>...", a.User)
    return args
}

