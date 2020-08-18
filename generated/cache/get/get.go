/*
## INFO
	Errors if the value doesn't exist.
## OPTIONS
	<key>
	: Cache key.
	[<group>]
	: Method for grouping data within the cache which allows the same key to be used across groups.
	---
	default: default
	---
## EXAMPLES
	    # Get cache.
	    $ wp cache get my_key my_group
	    my_value
	
 */
package get
import utils "github.com/bukowa/gowpcli"

// Get //Gets a value from the object cache.
type Get struct {
    Key string // <key>
    Group string // [<group>]
}

func (g Get) Args() []string {
    var args = []string{"cache", "get"}
    args = utils.MakeArg(args, "<key>", g.Key)
    args = utils.MakeArg(args, "[<group>]", g.Group)
    return args
}

