/*
## OPTIONS
	<user>...
	: One or more user IDs, user emails, or user logins.
## EXAMPLES
	    $ wp super-admin remove superadmin2
	    Success: Revoked super-admin capabilities.
	
 */
package remove
import utils "github.com/bukowa/gowpcli"

// Remove //Removes super admin privileges from one or more users.
type Remove struct {
    User []string // <user>...
}

func (r Remove) Args() []string {
    var args = []string{"super-admin", "remove"}
    args = utils.MakeArg(args, "<user>...", r.User)
    return args
}

