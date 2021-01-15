/*
## OPTIONS
	<user>
	: User ID, user email, or user login.
	<cap>
	: The capability to be removed.
## EXAMPLES
	    $ wp user remove-cap 11 publish_newsletters
	    Success: Removed 'publish_newsletters' cap for supervisor (11).
	    $ wp user remove-cap 11 publish_posts
	    Error: The 'publish_posts' cap for supervisor (11) is inherited from a role.
	    $ wp user remove-cap 11 nonexistent_cap
	    Error: No such 'nonexistent_cap' cap for supervisor (11).
	
 */
package removecap
import utils "github.com/bukowa/gowpcli"

// RemoveCap //Removes a user's capability.
type RemoveCap struct {
    User string // <user>
    Cap string // <cap>
}

func (r RemoveCap) Args() []string {
    var args = []string{"user", "remove-cap"}
    args = utils.MakeArg(args, "<user>", r.User)
    args = utils.MakeArg(args, "<cap>", r.Cap)
    return args
}

