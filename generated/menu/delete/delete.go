/*
## OPTIONS
	<menu>...
	: The name, slug, or term ID for the menu(s).
## EXAMPLES
	    $ wp menu delete "My Menu"
	    Success: 1 menu deleted.
	
 */
package delete
import utils "github.com/bukowa/gowpcli"

// Delete //Deletes one or more menus.
type Delete struct {
    Menu []string // <menu>...
}

func (d Delete) Args() []string {
    var args = []string{"menu", "delete"}
    args = utils.MakeArg(args, "<menu>...", d.Menu)
    return args
}

