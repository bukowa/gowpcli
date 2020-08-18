/*
## OPTIONS
	<id>
	: The ID of the object.
	<key>
	: The name of the meta field to update.
	[<value>]
	: The new value. If omitted, the value is read from STDIN.
	[--format=<format>]
	: The serialization format for the value.
	---
	default: plaintext
	options:
	  - plaintext
	  - json
	---
	
 */
package update
import utils "github.com/bukowa/gowpcli"

// Update //Update a meta field.
type Update struct {
    Id string // <id>
    Key string // <key>
    Value string // [<value>]
    Format string // [--format=<format>]
}

func (u Update) Args() []string {
    var args = []string{"term", "meta", "update"}
    args = utils.MakeArg(args, "<id>", u.Id)
    args = utils.MakeArg(args, "<key>", u.Key)
    args = utils.MakeArg(args, "[<value>]", u.Value)
    args = utils.MakeArg(args, "[--format=<format>]", u.Format)
    return args
}

