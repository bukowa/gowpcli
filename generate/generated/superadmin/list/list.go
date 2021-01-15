/*
## OPTIONS
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
	    # List user with super-admin capabilities
	    $ wp super-admin list
	    supervisor
	    administrator
	
 */
package list
import utils "github.com/bukowa/gowpcli"

// List //Lists users with super admin capabilities.
type List struct {
    Format string // [--format=<format>]
}

func (l List) Args() []string {
    var args = []string{"super-admin", "list"}
    args = utils.MakeArg(args, "[--format=<format>]", l.Format)
    return args
}

