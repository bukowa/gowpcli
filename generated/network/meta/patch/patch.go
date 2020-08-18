/*
## OPTIONS
	<action>
	: Patch action to perform.
	---
	options:
	  - insert
	  - update
	  - delete
	---
	<id>
	: The ID of the object.
	<key>
	: The name of the meta field to update.
	<key-path>...
	: The name(s) of the keys within the value to locate the value to patch.
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
package patch
import utils "github.com/bukowa/gowpcli"

// Patch //Update a nested value for a meta field.
type Patch struct {
    Action string // <action>
    Id string // <id>
    Key string // <key>
    KeyPath []string // <key-path>...
    Value string // [<value>]
    Format string // [--format=<format>]
}

func (p Patch) Args() []string {
    var args = []string{"network", "meta", "patch"}
    args = utils.MakeArg(args, "<action>", p.Action)
    args = utils.MakeArg(args, "<id>", p.Id)
    args = utils.MakeArg(args, "<key>", p.Key)
    args = utils.MakeArg(args, "<key-path>...", p.KeyPath)
    args = utils.MakeArg(args, "[<value>]", p.Value)
    args = utils.MakeArg(args, "[--format=<format>]", p.Format)
    return args
}

