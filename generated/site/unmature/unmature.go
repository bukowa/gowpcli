/*
## OPTIONS
	<id>...
	: One or more IDs of sites to set as unmature.
## EXAMPLES
	    $ wp site general 123
	    Success: Site 123 marked as unmature.
	
 */
package unmature
import utils "github.com/bukowa/gowpcli"

// Unmature //Sets one or more sites as immature.
type Unmature struct {
    Id []string // <id>...
}

func (u Unmature) Args() []string {
    var args = []string{"site", "unmature"}
    args = utils.MakeArg(args, "<id>...", u.Id)
    return args
}

