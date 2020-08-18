/*
## OPTIONS
	<role>
	: Key for the role.
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
	[--show-grant]
	: Display all capabilities defined for a role including grant.
	---
	default: false
	---
## EXAMPLES
	    # Display alphabetical list of Contributor capabilities.
	    $ wp cap list 'contributor' | sort
	    delete_posts
	    edit_posts
	    level_0
	    level_1
	    read
	
 */
package list
import utils "github.com/bukowa/gowpcli"

// List //Lists capabilities for a given role.
type List struct {
    Role string // <role>
    Format string // [--format=<format>]
    ShowGrant bool // [--show-grant]
}

func (l List) Args() []string {
    var args = []string{"cap", "list"}
    args = utils.MakeArg(args, "<role>", l.Role)
    args = utils.MakeArg(args, "[--format=<format>]", l.Format)
    args = utils.MakeArg(args, "[--show-grant]", l.ShowGrant)
    return args
}

