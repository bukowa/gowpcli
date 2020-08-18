/*
## OPTIONS
	<user>
	: The user login, user email, or user ID of the user to delete metadata from.
	<key>
	: The metadata key.
	[<value>]
	: The value to delete. If omitted, all rows with key will deleted.
## EXAMPLES
	    # Delete user meta
	    $ wp user meta delete 123 bio
	    Success: Deleted custom field.
	
 */
package delete
import utils "github.com/bukowa/gowpcli"

// Delete //Deletes a meta field.
type Delete struct {
    User string // <user>
    Key string // <key>
    Value string // [<value>]
}

func (d Delete) Args() []string {
    var args = []string{"user", "meta", "delete"}
    args = utils.MakeArg(args, "<user>", d.User)
    args = utils.MakeArg(args, "<key>", d.Key)
    args = utils.MakeArg(args, "[<value>]", d.Value)
    return args
}

