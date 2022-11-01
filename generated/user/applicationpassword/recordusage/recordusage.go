/*
## OPTIONS
	<user>
	: The user login, user email, or user ID of the user to update the application password for.
	<uuid>
	: The universally unique ID of the application password.
## EXAMPLES
	    # Record usage of an application password
	    $ wp user application-password record-usage 123 6633824d-c1d7-4f79-9dd5-4586f734d69e
	    Success: Recorded application password usage.
	
 */
package recordusage
import utils "github.com/bukowa/gowpcli"

// RecordUsage //Record usage of an application password.
type RecordUsage struct {
    User string // <user>
    Uuid string // <uuid>
}

func (r RecordUsage) Args() []string {
    var args = []string{"user", "application-password", "record-usage"}
    args = utils.MakeArg(args, "<user>", r.User)
    args = utils.MakeArg(args, "<uuid>", r.Uuid)
    return args
}

