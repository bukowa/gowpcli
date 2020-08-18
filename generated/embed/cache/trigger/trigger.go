/*
## INFO
	oEmbed caches for a post are stored in the post's metadata.
## OPTIONS
	<post_id>
	: ID of the post to do the caching for.
## EXAMPLES
	    # Triggers cache for a post
	    $ wp embed cache trigger 456
	    Success: Caching triggered!
	
 */
package trigger
import utils "github.com/bukowa/gowpcli"

// Trigger //Triggers the caching of all oEmbed results for a given post.
type Trigger struct {
    PostId string // <post_id>
}

func (t Trigger) Args() []string {
    var args = []string{"embed", "cache", "trigger"}
    args = utils.MakeArg(args, "<post_id>", t.PostId)
    return args
}

