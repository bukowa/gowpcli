/*
## OPTIONS
	<menu>
	: The name, slug, or term ID for the menu.
	<location>
	: Location's slug.
## EXAMPLES
	    $ wp menu location remove primary-menu primary
	    Success: Removed location from menu.
	
 */
package remove
import utils "github.com/bukowa/gowpcli"

// Remove //Removes a location from a menu.
type Remove struct {
    Menu string // <menu>
    Location string // <location>
}

func (r Remove) Args() []string {
    var args = []string{"menu", "location", "remove"}
    args = utils.MakeArg(args, "<menu>", r.Menu)
    args = utils.MakeArg(args, "<location>", r.Location)
    return args
}

