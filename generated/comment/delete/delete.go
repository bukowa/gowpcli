/*
## OPTIONS
	<id>...
	: One or more IDs of comments to delete.
	[--force]
	: Skip the trash bin.
## EXAMPLES
	    # Delete comment.
	    $ wp comment delete 1337 --force
	    Success: Deleted comment 1337.
	    # Delete multiple comments.
	    $ wp comment delete 1337 2341 --force
	    Success: Deleted comment 1337.
	    Success: Deleted comment 2341.
	
 */
package delete
import utils "github.com/bukowa/gowpcli"

// Delete //Deletes a comment.
type Delete struct {
    Id []string // <id>...
    Force bool // [--force]
}

func (d Delete) Args() []string {
    var args = []string{"comment", "delete"}
    args = utils.MakeArg(args, "<id>...", d.Id)
    args = utils.MakeArg(args, "[--force]", d.Force)
    return args
}

