/*
## OPTIONS
	<id>...
	: One or more IDs of sites to activate.
## EXAMPLES
	    $ wp site activate 123
	    Success: Site 123 activated.
	
 */
package activate
import utils "github.com/bukowa/gowpcli"

// Activate //Activates one or more sites.
type Activate struct {
    Id []string // <id>...
}

func (a Activate) Args() []string {
    var args = []string{"site", "activate"}
    args = utils.MakeArg(args, "<id>...", a.Id)
    return args
}

