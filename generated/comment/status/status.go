/*
## OPTIONS
	<id>
	: The ID of the comment to check.
## EXAMPLES
	    # Get status of comment.
	    $ wp comment status 1337
	    approved
	
 */
package status
import utils "github.com/bukowa/gowpcli"

// Status //Gets the status of a comment.
type Status struct {
    Id string // <id>
}

func (s Status) Args() []string {
    var args = []string{"comment", "status"}
    args = utils.MakeArg(args, "<id>", s.Id)
    return args
}

