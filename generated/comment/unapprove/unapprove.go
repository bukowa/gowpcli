/*
## OPTIONS
	<id>...
	: The IDs of the comments to unapprove.
## EXAMPLES
	    # Unapprove comment.
	    $ wp comment unapprove 1337
	    Success: Unapproved comment 1337.
	
 */
package unapprove
import utils "github.com/bukowa/gowpcli"

// Unapprove //Unapproves a comment.
type Unapprove struct {
    Id []string // <id>...
}

func (u Unapprove) Args() []string {
    var args = []string{"comment", "unapprove"}
    args = utils.MakeArg(args, "<id>...", u.Id)
    return args
}

