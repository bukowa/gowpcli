/*
## OPTIONS
	<id>...
	: The IDs of the comments to untrash.
## EXAMPLES
	    # Untrash comment.
	    $ wp comment untrash 1337
	    Success: Untrashed comment 1337.
	
 */
package untrash
import utils "github.com/bukowa/gowpcli"

// Untrash //Untrashes a comment.
type Untrash struct {
    Id []string // <id>...
}

func (u Untrash) Args() []string {
    var args = []string{"comment", "untrash"}
    args = utils.MakeArg(args, "<id>...", u.Id)
    return args
}

