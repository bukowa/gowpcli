/*
## OPTIONS
	<key>
	: Key for the alias.
## EXAMPLES
	    # Get alias.
	    $ wp cli alias get @prod
	    ssh: dev@somedeve.env:12345/home/dev/
	
 */
package get
import utils "github.com/bukowa/gowpcli"

// Get //Gets the value for an alias.
type Get struct {
    Key string // <key>
}

func (g Get) Args() []string {
    var args = []string{"cli", "alias", "get"}
    args = utils.MakeArg(args, "<key>", g.Key)
    return args
}

