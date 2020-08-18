/*
## OPTIONS
	<id>...
	: IDs for one or more posts to update.
## EXAMPLES
	    # Recount comment for the post.
	    $ wp comment recount 123
	    Updated post 123 comment count to 67.
	
 */
package recount
import utils "github.com/bukowa/gowpcli"

// Recount //Recalculates the comment_count value for one or more posts.
type Recount struct {
    Id []string // <id>...
}

func (r Recount) Args() []string {
    var args = []string{"comment", "recount"}
    args = utils.MakeArg(args, "<id>...", r.Id)
    return args
}

