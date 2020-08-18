/*
## OPTIONS
	<id>...
	: The IDs of the comments to mark as spam.
## EXAMPLES
	    # Spam comment.
	    $ wp comment spam 1337
	    Success: Marked as spam comment 1337.
	
 */
package spam
import utils "github.com/bukowa/gowpcli"

// Spam //Marks a comment as spam.
type Spam struct {
    Id []string // <id>...
}

func (s Spam) Args() []string {
    var args = []string{"comment", "spam"}
    args = utils.MakeArg(args, "<id>...", s.Id)
    return args
}

