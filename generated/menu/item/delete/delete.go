/*
## OPTIONS
	<db-id>...
	: Database ID for the menu item(s).
## EXAMPLES
	    $ wp menu item delete 45
	    Success: 1 menu item deleted.
	
 */
package delete
import utils "github.com/bukowa/gowpcli"

// Delete //Deletes one or more items from a menu.
type Delete struct {
    DbId []string // <db-id>...
}

func (d Delete) Args() []string {
    var args = []string{"menu", "item", "delete"}
    args = utils.MakeArg(args, "<db-id>...", d.DbId)
    return args
}

