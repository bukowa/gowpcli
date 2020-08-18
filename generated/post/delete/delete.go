/*
## OPTIONS
	<id>...
	: One or more IDs of posts to delete.
	[--force]
	: Skip the trash bin.
	[--defer-term-counting]
	: Recalculate term count in batch, for a performance boost.
## EXAMPLES
	    # Delete post skipping trash
	    $ wp post delete 123 --force
	    Success: Deleted post 123.
	    # Delete all pages
	    $ wp post delete $(wp post list --post_type='page' --format=ids)
	    Success: Trashed post 1164.
	    Success: Trashed post 1186.
	    # Delete all posts in the trash
	    $ wp post delete $(wp post list --post_status=trash --format=ids)
	    Success: Deleted post 1268.
	    Success: Deleted post 1294.
	
 */
package delete
import utils "github.com/bukowa/gowpcli"

// Delete //Deletes an existing post.
type Delete struct {
    Id []string // <id>...
    Force bool // [--force]
    DeferTermCounting bool // [--defer-term-counting]
}

func (d Delete) Args() []string {
    var args = []string{"post", "delete"}
    args = utils.MakeArg(args, "<id>...", d.Id)
    args = utils.MakeArg(args, "[--force]", d.Force)
    args = utils.MakeArg(args, "[--defer-term-counting]", d.DeferTermCounting)
    return args
}

