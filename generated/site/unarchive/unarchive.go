/*
## OPTIONS
	<id>...
	: One or more IDs of sites to unarchive.
## EXAMPLES
	    $ wp site unarchive 123
	    Success: Site 123 unarchived.
	
 */
package unarchive
import utils "github.com/bukowa/gowpcli"

// Unarchive //Unarchives one or more sites.
type Unarchive struct {
    Id []string // <id>...
}

func (u Unarchive) Args() []string {
    var args = []string{"site", "unarchive"}
    args = utils.MakeArg(args, "<id>...", u.Id)
    return args
}

