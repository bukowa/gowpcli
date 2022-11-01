/*
## OPTIONS
	<user>
	: The user login, user email, or user ID of the user to delete the application password for.
	[<uuid>...]
	: Comma-separated list of UUIDs of the application passwords to delete.
	[--all]
	: Delete all of the user's application password.
## EXAMPLES
	    # Record usage of an application password
	    $ wp user application-password record-usage 123 6633824d-c1d7-4f79-9dd5-4586f734d69e
	    Success: Recorded application password usage.
	
 */
package delete
import utils "github.com/bukowa/gowpcli"

// Delete //Delete an existing application password.
type Delete struct {
    User string // <user>
    Uuid []string // [<uuid>...]
    All bool // [--all]
}

func (d Delete) Args() []string {
    var args = []string{"user", "application-password", "delete"}
    args = utils.MakeArg(args, "<user>", d.User)
    args = utils.MakeArg(args, "[<uuid>...]", d.Uuid)
    args = utils.MakeArg(args, "[--all]", d.All)
    return args
}

