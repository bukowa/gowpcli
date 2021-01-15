/*
## OPTIONS
	<user>
	: User ID, user email, or user login.
	<cap>
	: The capability to add.
## EXAMPLES
	    # Add a capability for a user
	    $ wp user add-cap john create_premium_item
	    Success: Added 'create_premium_item' capability for john (16).
	    # Add a capability for a user
	    $ wp user add-cap 15 edit_product
	    Success: Added 'edit_product' capability for johndoe (15).
	
 */
package addcap
import utils "github.com/bukowa/gowpcli"

// AddCap //Adds a capability to a user.
type AddCap struct {
    User string // <user>
    Cap string // <cap>
}

func (a AddCap) Args() []string {
    var args = []string{"user", "add-cap"}
    args = utils.MakeArg(args, "<user>", a.User)
    args = utils.MakeArg(args, "<cap>", a.Cap)
    return args
}

