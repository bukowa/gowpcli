/*
## OPTIONS
	<id>...
	: One or more IDs of sites to archive.
## EXAMPLES
	    $ wp site archive 123
	    Success: Site 123 archived.
	
 */
package archive
import utils "github.com/bukowa/gowpcli"

// Archive //Archives one or more sites.
type Archive struct {
    Id []string // <id>...
}

func (a Archive) Args() []string {
    var args = []string{"site", "archive"}
    args = utils.MakeArg(args, "<id>...", a.Id)
    return args
}

