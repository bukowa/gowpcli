/*
## OPTIONS
	<user>
	: User ID, user email, or login.
	[--format=<format>]
	: Render output in a particular format.
	---
	default: list
	options:
	  - list
	  - table
	  - csv
	  - json
	  - count
	  - yaml
	---
## EXAMPLES
	    $ wp user list-caps 21
	    edit_product
	    create_premium_item
	
 */
package listcaps
import utils "github.com/bukowa/gowpcli"

// ListCaps //Lists all capabilities for a user.
type ListCaps struct {
    User string // <user>
    Format string // [--format=<format>]
}

func (l ListCaps) Args() []string {
    var args = []string{"user", "list-caps"}
    args = utils.MakeArg(args, "<user>", l.User)
    args = utils.MakeArg(args, "[--format=<format>]", l.Format)
    return args
}

