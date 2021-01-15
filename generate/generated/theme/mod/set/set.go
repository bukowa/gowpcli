/*
## OPTIONS
	<mod>
	: The name of the theme mod to set or update.
	<value>
	: The new value.
## EXAMPLES
	    # Set theme mod
	    $ wp theme mod set background_color 000000
	    Success: Theme mod background_color set to 000000
	
 */
package set
import utils "github.com/bukowa/gowpcli"

// Set //Sets the value of a theme mod.
type Set struct {
    Mod string // <mod>
    Value string // <value>
}

func (s Set) Args() []string {
    var args = []string{"theme", "mod", "set"}
    args = utils.MakeArg(args, "<mod>", s.Mod)
    args = utils.MakeArg(args, "<value>", s.Value)
    return args
}

