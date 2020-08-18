/*
## OPTIONS
	<user>
	: The user login, user email, or user ID of the user to update metadata for.
	<key>
	: The metadata key.
	<value>
	: The new metadata value.
	[--format=<format>]
	: The serialization format for the value.
	---
	default: plaintext
	options:
	  - plaintext
	  - json
	---
## EXAMPLES
	    # Update user meta
	    $ wp user meta update 123 bio "Mary is an awesome WordPress developer."
	    Success: Updated custom field 'bio'.
	
 */
package update
import utils "github.com/bukowa/gowpcli"

// Update //Updates a meta field.
type Update struct {
    User string // <user>
    Key string // <key>
    Value string // <value>
    Format string // [--format=<format>]
}

func (u Update) Args() []string {
    var args = []string{"user", "meta", "update"}
    args = utils.MakeArg(args, "<user>", u.User)
    args = utils.MakeArg(args, "<key>", u.Key)
    args = utils.MakeArg(args, "<value>", u.Value)
    args = utils.MakeArg(args, "[--format=<format>]", u.Format)
    return args
}

