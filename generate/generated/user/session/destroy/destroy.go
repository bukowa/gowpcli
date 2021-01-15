/*
## OPTIONS
	<user>
	: User ID, user email, or user login.
	[<token>]
	: The token of the session to destroy. Defaults to the most recently created session.
	[--all]
	: Destroy all of the user's sessions.
## EXAMPLES
	    # Destroy the most recent session of the given user.
	    $ wp user session destroy admin
	    Success: Destroyed session. 3 sessions remaining.
	    # Destroy a specific session of the given user.
	    $ wp user session destroy admin e073ad8540a9c2...
	    Success: Destroyed session. 2 sessions remaining.
	    # Destroy all the sessions of the given user.
	    $ wp user session destroy admin --all
	    Success: Destroyed all sessions.
	    # Destroy all sessions for all users.
	    $ wp user list --field=ID | xargs -n 1 wp user session destroy --all
	    Success: Destroyed all sessions.
	    Success: Destroyed all sessions.
	
 */
package destroy
import utils "github.com/bukowa/gowpcli"

// Destroy //Destroy a session for the given user.
type Destroy struct {
    User string // <user>
    Token string // [<token>]
    All bool // [--all]
}

func (d Destroy) Args() []string {
    var args = []string{"user", "session", "destroy"}
    args = utils.MakeArg(args, "<user>", d.User)
    args = utils.MakeArg(args, "[<token>]", d.Token)
    args = utils.MakeArg(args, "[--all]", d.All)
    return args
}

