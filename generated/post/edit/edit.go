/*
## OPTIONS
	<id>
	: The ID of the post to edit.
## EXAMPLES
	    # Launch system editor to edit post
	    $ wp post edit 123
	
 */
package edit
import utils "github.com/bukowa/gowpcli"

// Edit //Launches system editor to edit post content.
type Edit struct {
    Id string // <id>
}

func (e Edit) Args() []string {
    var args = []string{"post", "edit"}
    args = utils.MakeArg(args, "<id>", e.Id)
    return args
}

