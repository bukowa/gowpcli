/*
## OPTIONS
	<menu>
	: The name, slug, or term ID for the menu.
	<location>
	: Location's slug.
## EXAMPLES
	    $ wp menu location assign primary-menu primary
	    Success: Assigned location primary to menu primary-menu.
	
 */
package assign
import utils "github.com/bukowa/gowpcli"

// Assign //Assigns a location to a menu.
type Assign struct {
    Menu string // <menu>
    Location string // <location>
}

func (a Assign) Args() []string {
    var args = []string{"menu", "location", "assign"}
    args = utils.MakeArg(args, "<menu>", a.Menu)
    args = utils.MakeArg(args, "<location>", a.Location)
    return args
}

