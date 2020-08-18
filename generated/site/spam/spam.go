/*
## OPTIONS
	<id>...
	: One or more IDs of sites to be marked as spam.
## EXAMPLES
	    $ wp site spam 123
	    Success: Site 123 marked as spam.
	
 */
package spam
import utils "github.com/bukowa/gowpcli"

// Spam //Marks one or more sites as spam.
type Spam struct {
    Id []string // <id>...
}

func (s Spam) Args() []string {
    var args = []string{"site", "spam"}
    args = utils.MakeArg(args, "<id>...", s.Id)
    return args
}

