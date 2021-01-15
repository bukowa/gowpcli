/*
## OPTIONS
	<id>...
	: One or more IDs of sites to remove from spam.
## EXAMPLES
	    $ wp site unspam 123
	    Success: Site 123 removed from spam.
	
 */
package unspam
import utils "github.com/bukowa/gowpcli"

// Unspam //Removes one or more sites from spam.
type Unspam struct {
    Id []string // <id>...
}

func (u Unspam) Args() []string {
    var args = []string{"site", "unspam"}
    args = utils.MakeArg(args, "<id>...", u.Id)
    return args
}

