/*
## INFO
	Errors if the value can't be decremented.
## OPTIONS
	<key>
	: Cache key.
	[<offset>]
	: The amount by which to decrement the item's value.
	---
	default: 1
	---
	[<group>]
	: Method for grouping data within the cache which allows the same key to be used across groups.
	---
	default: default
	---
## EXAMPLES
	    # Decrease cache value.
	    $ wp cache decr my_key 2 my_group
	    48
	
 */
package decr
import utils "github.com/bukowa/gowpcli"

// Decr //Decrements a value in the object cache.
type Decr struct {
    Key string // <key>
    Offset string // [<offset>]
    Group string // [<group>]
}

func (d Decr) Args() []string {
    var args = []string{"cache", "decr"}
    args = utils.MakeArg(args, "<key>", d.Key)
    args = utils.MakeArg(args, "[<offset>]", d.Offset)
    args = utils.MakeArg(args, "[<group>]", d.Group)
    return args
}

