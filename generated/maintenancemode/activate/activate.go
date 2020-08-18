/*
## INFO
	[--force]
	: Force maintenance mode activation operation.
## EXAMPLES
	    $ wp maintenance-mode activate
	    Enabling Maintenance mode...
	    Success: Activated Maintenance mode.
	
 */
package activate
import utils "github.com/bukowa/gowpcli"

// Activate //Activates maintenance mode.
type Activate struct {
    Force bool // [--force]
}

func (a Activate) Args() []string {
    var args = []string{"maintenance-mode", "activate"}
    args = utils.MakeArg(args, "[--force]", a.Force)
    return args
}

