/*
## INFO
	Displays a success message if the comment does exist.
## OPTIONS
	<id>
	: The ID of the comment to check.
## EXAMPLES
	    # Check whether comment exists.
	    $ wp comment exists 1337
	    Success: Comment with ID 1337 exists.
	
 */
package exists
import utils "github.com/bukowa/gowpcli"

// Exists //Verifies whether a comment exists.
type Exists struct {
    Id string // <id>
}

func (e Exists) Args() []string {
    var args = []string{"comment", "exists"}
    args = utils.MakeArg(args, "<id>", e.Id)
    return args
}

