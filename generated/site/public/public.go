/*
## OPTIONS
	<id>...
	: One or more IDs of sites to set as public.
## EXAMPLES
	    $ wp site public 123
	    Success: Site 123 marked as public.
	
 */
package public
import utils "github.com/bukowa/gowpcli"

// Public //Sets one or more sites as public.
type Public struct {
    Id []string // <id>...
}

func (p Public) Args() []string {
    var args = []string{"site", "public"}
    args = utils.MakeArg(args, "<id>...", p.Id)
    return args
}

