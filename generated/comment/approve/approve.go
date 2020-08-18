/*
## OPTIONS
	<id>...
	: The IDs of the comments to approve.
## EXAMPLES
	    # Approve comment.
	    $ wp comment approve 1337
	    Success: Approved comment 1337.
	
 */
package approve
import utils "github.com/bukowa/gowpcli"

// Approve //Approves a comment.
type Approve struct {
    Id []string // <id>...
}

func (a Approve) Args() []string {
    var args = []string{"comment", "approve"}
    args = utils.MakeArg(args, "<id>...", a.Id)
    return args
}

