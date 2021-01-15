/*
## OPTIONS
	<user>...
	: one or more user logins or IDs.
	[--skip-email]
	: Don't send an email notification to the affected user(s).
## EXAMPLES
	    # Reset the password for two users and send them the change email.
	    $ wp user reset-password admin editor
	    Reset password for admin.
	    Reset password for editor.
	    Success: Passwords reset for 2 users.
	
 */
package resetpassword
import utils "github.com/bukowa/gowpcli"

// ResetPassword //Resets the password for one or more users.
type ResetPassword struct {
    User []string // <user>...
    SkipEmail bool // [--skip-email]
}

func (r ResetPassword) Args() []string {
    var args = []string{"user", "reset-password"}
    args = utils.MakeArg(args, "<user>...", r.User)
    args = utils.MakeArg(args, "[--skip-email]", r.SkipEmail)
    return args
}

