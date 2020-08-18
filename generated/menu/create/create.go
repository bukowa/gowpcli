/*
## OPTIONS
	<menu-name>
	: A descriptive name for the menu.
	[--porcelain]
	: Output just the new menu id.
## EXAMPLES
	    $ wp menu create "My Menu"
	    Success: Created menu 200.
	
 */
package create
import utils "github.com/bukowa/gowpcli"

// Create //Creates a new menu.
type Create struct {
    MenuName string // <menu-name>
    Porcelain bool // [--porcelain]
}

func (c Create) Args() []string {
    var args = []string{"menu", "create"}
    args = utils.MakeArg(args, "<menu-name>", c.MenuName)
    args = utils.MakeArg(args, "[--porcelain]", c.Porcelain)
    return args
}

