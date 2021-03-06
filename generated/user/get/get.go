/*
## OPTIONS
	<user>
	: User ID, user email, or user login.
	[--field=<field>]
	: Instead of returning the whole user, returns the value of a single field.
	[--fields=<fields>]
	: Get a specific subset of the user's fields.
	[--format=<format>]
	: Render output in a particular format.
	---
	default: table
	options:
	  - table
	  - csv
	  - json
	  - yaml
	---
## EXAMPLES
	    # Get user
	    $ wp user get 12 --field=login
	    supervisor
	    # Get user and export to JSON file
	    $ wp user get bob --format=json > bob.json
	
 */
package get
import utils "github.com/bukowa/gowpcli"

// Get //Gets details about a user.
type Get struct {
    User string // <user>
    Field string // [--field=<field>]
    Fields string // [--fields=<fields>]
    Format string // [--format=<format>]
}

func (g Get) Args() []string {
    var args = []string{"user", "get"}
    args = utils.MakeArg(args, "<user>", g.User)
    args = utils.MakeArg(args, "[--field=<field>]", g.Field)
    args = utils.MakeArg(args, "[--fields=<fields>]", g.Fields)
    args = utils.MakeArg(args, "[--format=<format>]", g.Format)
    return args
}

