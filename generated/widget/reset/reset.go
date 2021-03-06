/*
## INFO
	Removes all widgets from the sidebar and places them in Inactive Widgets.
## OPTIONS
	[<sidebar-id>...]
	: One or more sidebars to reset.
	[--all]
	: If set, all sidebars will be reset.
## EXAMPLES
	    # Reset a sidebar
	    $ wp widget reset sidebar-1
	    Success: Sidebar 'sidebar-1' reset.
	    # Reset multiple sidebars
	    $ wp widget reset sidebar-1 sidebar-2
	    Success: Sidebar 'sidebar-1' reset.
	    Success: Sidebar 'sidebar-2' reset.
	    # Reset all sidebars
	    $ wp widget reset --all
	    Success: Sidebar 'sidebar-1' reset.
	    Success: Sidebar 'sidebar-2' reset.
	    Success: Sidebar 'sidebar-3' reset.
	
 */
package reset
import utils "github.com/bukowa/gowpcli"

// Reset //Resets sidebar.
type Reset struct {
    SidebarId []string // [<sidebar-id>...]
    All bool // [--all]
}

func (r Reset) Args() []string {
    var args = []string{"widget", "reset"}
    args = utils.MakeArg(args, "[<sidebar-id>...]", r.SidebarId)
    args = utils.MakeArg(args, "[--all]", r.All)
    return args
}

