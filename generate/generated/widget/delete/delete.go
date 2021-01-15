/*
## OPTIONS
	<widget-id>...
	: Unique ID for the widget(s)
## EXAMPLES
	    # Delete the recent-comments-2 widget from its sidebar.
	    $ wp widget delete recent-comments-2
	    Success: Deleted 1 of 1 widgets.
	
 */
package delete
import utils "github.com/bukowa/gowpcli"

// Delete //Deletes one or more widgets from a sidebar.
type Delete struct {
    WidgetId []string // <widget-id>...
}

func (d Delete) Args() []string {
    var args = []string{"widget", "delete"}
    args = utils.MakeArg(args, "<widget-id>...", d.WidgetId)
    return args
}

