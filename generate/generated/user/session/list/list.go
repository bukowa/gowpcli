/*
## INFO
	Note: The `token` field does not return the actual token, but a hash of
	it. The real token is not persisted and can only be found in the
	corresponding cookies on the client side.
## OPTIONS
	<user>
	: User ID, user email, or user login.
	[--fields=<fields>]
	: Limit the output to specific fields.
	[--format=<format>]
	: Render output in a particular format.
	---
	default: table
	options:
	  - table
	  - csv
	  - json
	  - yaml
	  - count
	  - ids
	---
## AVAILABLE FIELDS
	These fields will be displayed by default for each session:
	* token
	* login_time
	* expiration_time
	* ip
	* ua
	These fields are optionally available:
	* expiration
	* login
## EXAMPLES
	    # List a user's sessions.
	    $ wp user session list admin@example.com --format=csv
	    login_time,expiration_time,ip,ua
	    "2016-01-01 12:34:56","2016-02-01 12:34:56",127.0.0.1,"Mozilla/5.0..."
	
 */
package list
import utils "github.com/bukowa/gowpcli"

// List //List sessions for the given user.
type List struct {
    User string // <user>
    Fields string // [--fields=<fields>]
    Format string // [--format=<format>]
}

func (l List) Args() []string {
    var args = []string{"user", "session", "list"}
    args = utils.MakeArg(args, "<user>", l.User)
    args = utils.MakeArg(args, "[--fields=<fields>]", l.Fields)
    args = utils.MakeArg(args, "[--format=<format>]", l.Format)
    return args
}

