/*
## OPTIONS
	<key>
	: The name of the site option to add.
	[<value>]
	: The value of the site option to add. If ommited, the value is read from STDIN.
	[--format=<format>]
	: The serialization format for the value.
	---
	default: plaintext
	options:
	  - plaintext
	  - json
	---
## EXAMPLES
	    # Create a site option by reading a JSON file
	    $ wp site option add my_option --format=json < config.json
	    Success: Added 'my_option' site option.
	
 */
package add
import utils "github.com/bukowa/gowpcli"

// Add //Adds a site option.
type Add struct {
    Key string // <key>
    Value string // [<value>]
    Format string // [--format=<format>]
}

func (a Add) Args() []string {
    var args = []string{"site", "option", "add"}
    args = utils.MakeArg(args, "<key>", a.Key)
    args = utils.MakeArg(args, "[<value>]", a.Value)
    args = utils.MakeArg(args, "[--format=<format>]", a.Format)
    return args
}

