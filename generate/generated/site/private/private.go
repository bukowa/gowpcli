/*
## OPTIONS
	<id>...
	: One or more IDs of sites to set as private.
## EXAMPLES
	    $ wp site private 123
	    Success: Site 123 marked as private.
	
 */
package private
import utils "github.com/bukowa/gowpcli"

// Private //Sets one or more sites as private.
type Private struct {
    Id []string // <id>...
}

func (p Private) Args() []string {
    var args = []string{"site", "private"}
    args = utils.MakeArg(args, "<id>...", p.Id)
    return args
}

