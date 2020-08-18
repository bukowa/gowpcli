/*
## OPTIONS
	<key>
	: The name of the site option to update.
	[<value>]
	: The new value. If ommited, the value is read from STDIN.
	[--format=<format>]
	: The serialization format for the value.
	---
	default: plaintext
	options:
	  - plaintext
	  - json
	---
## EXAMPLES
	    # Update a site option by reading from a file
	    $ wp site option update my_option < value.txt
	    Success: Updated 'my_option' site option.
	
 */
package update
import utils "github.com/bukowa/gowpcli"

// Update //Updates a site option.
type Update struct {
    Key string // <key>
    Value string // [<value>]
    Format string // [--format=<format>]
}

func (u Update) Args() []string {
    var args = []string{"site", "option", "update"}
    args = utils.MakeArg(args, "<key>", u.Key)
    args = utils.MakeArg(args, "[<value>]", u.Value)
    args = utils.MakeArg(args, "[--format=<format>]", u.Format)
    return args
}

