/*
## OPTIONS
	<key>
	: The option name.
	<key-path>...
	: The name(s) of the keys within the value to locate the value to pluck.
	[--format=<format>]
	: The output format of the value.
	---
	default: plaintext
	options:
	  - plaintext
	  - json
	  - yaml
	---
	
 */
package pluck
import utils "github.com/bukowa/gowpcli"

// Pluck //Gets a nested value from an option.
type Pluck struct {
    Key string // <key>
    KeyPath []string // <key-path>...
    Format string // [--format=<format>]
}

func (p Pluck) Args() []string {
    var args = []string{"option", "pluck"}
    args = utils.MakeArg(args, "<key>", p.Key)
    args = utils.MakeArg(args, "<key-path>...", p.KeyPath)
    args = utils.MakeArg(args, "[--format=<format>]", p.Format)
    return args
}

