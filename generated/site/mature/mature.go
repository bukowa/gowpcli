/*
## OPTIONS
	<id>...
	: One or more IDs of sites to set as mature.
## EXAMPLES
	    $ wp site mature 123
	    Success: Site 123 marked as mature.
	
 */
package mature
import utils "github.com/bukowa/gowpcli"

// Mature //Sets one or more sites as mature.
type Mature struct {
    Id []string // <id>...
}

func (m Mature) Args() []string {
    var args = []string{"site", "mature"}
    args = utils.MakeArg(args, "<id>...", m.Id)
    return args
}

