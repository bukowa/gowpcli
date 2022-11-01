/*
## OPTIONS
	<id>
	: ID of the post to clean.
	<key>
	: Meta key to clean up.
## EXAMPLES
	    # Delete duplicate post meta.
	    wp post meta clean-duplicates 1234 enclosure
	    Success: Cleaned up duplicate 'enclosure' meta values.
	
 */
package cleanduplicates
import utils "github.com/bukowa/gowpcli"

// CleanDuplicates //Cleans up duplicate post meta values on a post.
type CleanDuplicates struct {
    Id string // <id>
    Key string // <key>
}

func (c CleanDuplicates) Args() []string {
    var args = []string{"post", "meta", "clean-duplicates"}
    args = utils.MakeArg(args, "<id>", c.Id)
    args = utils.MakeArg(args, "<key>", c.Key)
    return args
}

