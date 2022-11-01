/*
## INFO
	Displays a success message if the post does exist.
## OPTIONS
	<id>
	: The ID of the post to check.
## EXAMPLES
	    # The post exists.
	    $ wp post exists 1337
	    Success: Post with ID 1337 exists.
	    $ echo $?
	    0
	    # The post does not exist.
	    $ wp post exists 10000
	    $ echo $?
	    1
	
 */
package exists
import utils "github.com/bukowa/gowpcli"

// Exists //Verifies whether a post exists.
type Exists struct {
    Id string // <id>
}

func (e Exists) Args() []string {
    var args = []string{"post", "exists"}
    args = utils.MakeArg(args, "<id>", e.Id)
    return args
}

