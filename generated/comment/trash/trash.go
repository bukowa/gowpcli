/*
## OPTIONS
	<id>...
	: The IDs of the comments to trash.
## EXAMPLES
	    # Trash comment.
	    $ wp comment trash 1337
	    Success: Trashed comment 1337.
	
 */
package trash
import utils "github.com/bukowa/gowpcli"

// Trash //Trashes a comment.
type Trash struct {
    Id []string // <id>...
}

func (t Trash) Args() []string {
    var args = []string{"comment", "trash"}
    args = utils.MakeArg(args, "<id>...", t.Id)
    return args
}

