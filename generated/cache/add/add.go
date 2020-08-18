/*
## INFO
	Errors if a value already exists for the key, which means the value can't
	be added.
## OPTIONS
	<key>
	: Cache key.
	<value>
	: Value to add to the key.
	[<group>]
	: Method for grouping data within the cache which allows the same key to be used across groups.
	---
	default: default
	---
	[<expiration>]
	: Define how long to keep the value, in seconds. `0` means as long as possible.
	---
	default: 0
	---
## EXAMPLES
	    # Add cache.
	    $ wp cache add my_key my_group my_value 300
	    Success: Added object 'my_key' in group 'my_value'.
	
 */
package add
import utils "github.com/bukowa/gowpcli"

// Add //Adds a value to the object cache.
type Add struct {
    Key string // <key>
    Value string // <value>
    Group string // [<group>]
    Expiration string // [<expiration>]
}

func (a Add) Args() []string {
    var args = []string{"cache", "add"}
    args = utils.MakeArg(args, "<key>", a.Key)
    args = utils.MakeArg(args, "<value>", a.Value)
    args = utils.MakeArg(args, "[<group>]", a.Group)
    args = utils.MakeArg(args, "[<expiration>]", a.Expiration)
    return args
}

