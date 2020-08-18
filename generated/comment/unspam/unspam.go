/*
## OPTIONS
	<id>...
	: The IDs of the comments to unmark as spam.
## EXAMPLES
	    # Unspam comment.
	    $ wp comment unspam 1337
	    Success: Unspammed comment 1337.
	
 */
package unspam
import utils "github.com/bukowa/gowpcli"

// Unspam //Unmarks a comment as spam.
type Unspam struct {
    Id []string // <id>...
}

func (u Unspam) Args() []string {
    var args = []string{"comment", "unspam"}
    args = utils.MakeArg(args, "<id>...", u.Id)
    return args
}

