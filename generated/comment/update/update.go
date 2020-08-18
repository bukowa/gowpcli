/*
## OPTIONS
	<id>...
	: One or more IDs of comments to update.
	--<field>=<value>
	: One or more fields to update. See wp_update_comment().
## EXAMPLES
	    # Update comment.
	    $ wp comment update 123 --comment_author='That Guy'
	    Success: Updated comment 123.
	
 */
package update
import utils "github.com/bukowa/gowpcli"

// Update //Updates one or more comments.
type Update struct {
    Id []string // <id>...
    Field string // --<field>=<value>
}

func (u Update) Args() []string {
    var args = []string{"comment", "update"}
    args = utils.MakeArg(args, "<id>...", u.Id)
    args = utils.MakeArg(args, "--<field>=<value>", u.Field)
    return args
}

