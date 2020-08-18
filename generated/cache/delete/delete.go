/*
## INFO
	Errors if the value can't be deleted.
## OPTIONS
	<key>
	: Cache key.
	[<group>]
	: Method for grouping data within the cache which allows the same key to be used across groups.
	---
	default: default
	---
## EXAMPLES
	    # Delete cache.
	    $ wp cache delete my_key my_group
	    Success: Object deleted.
	
 */
package delete
import utils "github.com/bukowa/gowpcli"

// Delete //Removes a value from the object cache.
type Delete struct {
    Key string // <key>
    Group string // [<group>]
}

func (d Delete) Args() []string {
    var args = []string{"cache", "delete"}
    args = utils.MakeArg(args, "<key>", d.Key)
    args = utils.MakeArg(args, "[<group>]", d.Group)
    return args
}

