/*
## OPTIONS
	<id>
	: The ID of the object.
	<key>
	: The name of the meta field to get.
	[--format=<format>]
	: Get value in a particular format.
	---
	default: var_export
	options:
	  - var_export
	  - json
	  - yaml
	---
	
 */
package get
import utils "github.com/bukowa/gowpcli"

// Get //Get meta field value.
type Get struct {
    Id string // <id>
    Key string // <key>
    Format string // [--format=<format>]
}

func (g Get) Args() []string {
    var args = []string{"post", "meta", "get"}
    args = utils.MakeArg(args, "<id>", g.Id)
    args = utils.MakeArg(args, "<key>", g.Key)
    args = utils.MakeArg(args, "[--format=<format>]", g.Format)
    return args
}

