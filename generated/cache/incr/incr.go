/*
## INFO
	Errors if the value can't be incremented.
## OPTIONS
	<key>
	: Cache key.
	[<offset>]
	: The amount by which to increment the item's value.
	---
	default: 1
	---
	[<group>]
	: Method for grouping data within the cache which allows the same key to be used across groups.
	---
	default: default
	---
## EXAMPLES
	    # Increase cache value.
	    $ wp cache incr my_key 2 my_group
	    50
	
 */
package incr
import utils "github.com/bukowa/gowpcli"

// Incr //Increments a value in the object cache.
type Incr struct {
    Key string // <key>
    Offset string // [<offset>]
    Group string // [<group>]
}

func (i Incr) Args() []string {
    var args = []string{"cache", "incr"}
    args = utils.MakeArg(args, "<key>", i.Key)
    args = utils.MakeArg(args, "[<offset>]", i.Offset)
    args = utils.MakeArg(args, "[<group>]", i.Group)
    return args
}

