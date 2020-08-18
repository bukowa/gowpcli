/*
## INFO
	oEmbed caches for a post are stored in the post's metadata.
## OPTIONS
	<post_id>
	: ID of the post to clear the cache for.
## EXAMPLES
	    # Clear cache for a post
	    $ wp embed cache clear 123
	    Success: Cleared oEmbed cache.
	
 */
package clear
import utils "github.com/bukowa/gowpcli"

// Clear //Deletes all oEmbed caches for a given post.
type Clear struct {
    PostId string // <post_id>
}

func (c Clear) Args() []string {
    var args = []string{"embed", "cache", "clear"}
    args = utils.MakeArg(args, "<post_id>", c.PostId)
    return args
}

