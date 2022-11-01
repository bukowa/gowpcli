/*
## INFO
	On multisite, `wp user delete` only removes the user from the current
	site. Include `--network` to also remove the user from the database, but
	make sure to reassign their posts prior to deleting the user.
## OPTIONS
	<user>...
	: The user login, user email, or user ID of the user(s) to delete.
	[--network]
	: On multisite, delete the user from the entire network.
	[--reassign=<user-id>]
	: User ID to reassign the posts to.
	[--yes]
	: Answer yes to any confirmation prompts.
## EXAMPLES
	    # Delete user 123 and reassign posts to user 567
	    $ wp user delete 123 --reassign=567
	    Success: Removed user 123 from http://example.com
	    # Delete all contributors and reassign their posts to user 2
	    $ wp user delete $(wp user list --role=contributor --field=ID) --reassign=2
	    Success: Removed user 813 from http://example.com
	    Success: Removed user 578 from http://example.com
	    # Delete all contributors in batches of 100 (avoid error: argument list too long: wp)
	    $ wp user delete $(wp user list --role=contributor --field=ID | head -n 100)
	
 */
package delete
import utils "github.com/bukowa/gowpcli"

// Delete //Deletes one or more users from the current site.
type Delete struct {
    User []string // <user>...
    Network bool // [--network]
    Reassign string // [--reassign=<user-id>]
    Yes bool // [--yes]
}

func (d Delete) Args() []string {
    var args = []string{"user", "delete"}
    args = utils.MakeArg(args, "<user>...", d.User)
    args = utils.MakeArg(args, "[--network]", d.Network)
    args = utils.MakeArg(args, "[--reassign=<user-id>]", d.Reassign)
    args = utils.MakeArg(args, "[--yes]", d.Yes)
    return args
}

