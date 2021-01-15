/*
## INFO
	Creates a specified number of new users with dummy data.
## OPTIONS
	[--count=<number>]
	: How many users to generate?
	---
	default: 100
	---
	[--role=<role>]
	: The role of the generated users. Default: default role from WP
	[--format=<format>]
	: Render output in a particular format.
	---
	default: progress
	options:
	  - progress
	  - ids
	---
## EXAMPLES
	    # Add meta to every generated users.
	    $ wp user generate --format=ids --count=3 | xargs -d ' ' -I % wp user meta add % foo bar
	    Success: Added custom field.
	    Success: Added custom field.
	    Success: Added custom field.
	
 */
package generate
import utils "github.com/bukowa/gowpcli"

// Generate //Generates some users.
type Generate struct {
    Count string // [--count=<number>]
    Role string // [--role=<role>]
    Format string // [--format=<format>]
}

func (g Generate) Args() []string {
    var args = []string{"user", "generate"}
    args = utils.MakeArg(args, "[--count=<number>]", g.Count)
    args = utils.MakeArg(args, "[--role=<role>]", g.Role)
    args = utils.MakeArg(args, "[--format=<format>]", g.Format)
    return args
}

