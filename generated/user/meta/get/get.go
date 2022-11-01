/*
## OPTIONS
	<user>
	: The user login, user email, or user ID of the user to get metadata for.
	<key>
	: The metadata key.
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
	    # Get user meta
	    $ wp user meta get 123 bio
	    Mary is an WordPress developer.
	    # Get the primary site of a user (for multisite)
	    $ wp user meta get 2 primary_blog
	    3
	
 */
package get
import utils "github.com/bukowa/gowpcli"

// Get //Gets meta field value.
type Get struct {
    User string // <user>
    Key string // <key>
    Format string // [--format=<format>]
}

func (g Get) Args() []string {
    var args = []string{"user", "meta", "get"}
    args = utils.MakeArg(args, "<user>", g.User)
    args = utils.MakeArg(args, "<key>", g.Key)
    args = utils.MakeArg(args, "[--format=<format>]", g.Format)
    return args
}

