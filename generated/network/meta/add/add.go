/*
## OPTIONS
	<id>
	: The ID of the object.
	<key>
	: The name of the meta field to create.
	[<value>]
	: The value of the meta field. If omitted, the value is read from STDIN.
	[--format=<format>]
	: The serialization format for the value.
	---
	default: plaintext
	options:
	  - plaintext
	  - json
	---
	
 */
package add
import utils "github.com/bukowa/gowpcli"

// Add //Add a meta field.
type Add struct {
    Id string // <id>
    Key string // <key>
    Value string // [<value>]
    Format string // [--format=<format>]
}

func (a Add) Args() []string {
    var args = []string{"network", "meta", "add"}
    args = utils.MakeArg(args, "<id>", a.Id)
    args = utils.MakeArg(args, "<key>", a.Key)
    args = utils.MakeArg(args, "[<value>]", a.Value)
    args = utils.MakeArg(args, "[--format=<format>]", a.Format)
    return args
}

