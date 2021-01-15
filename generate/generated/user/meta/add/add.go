/*
## OPTIONS
	<user>
	: The user login, user email, or user ID of the user to add metadata for.
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
	    # Add user meta
	    $ wp user meta add 123 bio "Mary is an WordPress developer."
	    Success: Added custom field.
	
 */
package add
import utils "github.com/bukowa/gowpcli"

// Add //Adds a meta field.
type Add struct {
    User string // <user>
    Key string // <key>
    Value string // <value>
    Format string // [--format=<format>]
}

func (a Add) Args() []string {
    var args = []string{"user", "meta", "add"}
    args = utils.MakeArg(args, "<user>", a.User)
    args = utils.MakeArg(args, "<key>", a.Key)
    args = utils.MakeArg(args, "<value>", a.Value)
    args = utils.MakeArg(args, "[--format=<format>]", a.Format)
    return args
}

