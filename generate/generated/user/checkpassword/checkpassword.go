/*
## OPTIONS
	<user>
	: The user login, user email or user ID of the user to check credentials for.
	<user_pass>
	: A string that contains the plain text password for the user.
## EXAMPLES
	    # Check whether given credentials are valid; exit status 0 if valid, otherwise 1
	    $ wp user check-password admin adminpass
	    $ echo $?
	    1
	    # Bash script for checking whether given credentials are valid or not
	    if ! $(wp user check-password admin adminpass); then
	     notify-send "Invalid Credentials";
	    fi
	
 */
package checkpassword
import utils "github.com/bukowa/gowpcli"

// CheckPassword //Checks if a user's password is valid or not.
type CheckPassword struct {
    User string // <user>
    UserPass string // <user_pass>
}

func (c CheckPassword) Args() []string {
    var args = []string{"user", "check-password"}
    args = utils.MakeArg(args, "<user>", c.User)
    args = utils.MakeArg(args, "<user_pass>", c.UserPass)
    return args
}

