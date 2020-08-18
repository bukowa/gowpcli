/*
## OPTIONS
	<id>...
	: One or more IDs of sites to deactivate.
## EXAMPLES
	    $ wp site deactivate 123
	    Success: Site 123 deactivated.
	
 */
package deactivate
import utils "github.com/bukowa/gowpcli"

// Deactivate //Deactivates one or more sites.
type Deactivate struct {
    Id []string // <id>...
}

func (d Deactivate) Args() []string {
    var args = []string{"site", "deactivate"}
    args = utils.MakeArg(args, "<id>...", d.Id)
    return args
}

