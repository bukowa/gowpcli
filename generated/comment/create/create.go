/*
## OPTIONS
	[--<field>=<value>]
	: Associative args for the new comment. See wp_insert_comment().
	[--porcelain]
	: Output just the new comment id.
## EXAMPLES
	    # Create comment.
	    $ wp comment create --comment_post_ID=15 --comment_content="hello blog" --comment_author="wp-cli"
	    Success: Created comment 932.
	
 */
package create
import utils "github.com/bukowa/gowpcli"

// Create //Creates a new comment.
type Create struct {
    FieldMap map[string]string // [--<field>=<value>]
    Porcelain bool // [--porcelain]
}

func (c Create) Args() []string {
    var args = []string{"comment", "create"}
    args = utils.MakeArg(args, "[--<field>=<value>]", c.FieldMap)
    args = utils.MakeArg(args, "[--porcelain]", c.Porcelain)
    return args
}

