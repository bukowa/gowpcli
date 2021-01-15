/*
## INFO
	Moves widgets to Inactive Widgets.
## OPTIONS
	<widget-id>...
	: Unique ID for the widget(s)
## EXAMPLES
	    # Deactivate the recent-comments-2 widget.
	    $ wp widget deactivate recent-comments-2
	    Success: 1 widget deactivated.
	
 */
package deactivate
import utils "github.com/bukowa/gowpcli"

// Deactivate //Deactivates one or more widgets from an active sidebar.
type Deactivate struct {
    WidgetId []string // <widget-id>...
}

func (d Deactivate) Args() []string {
    var args = []string{"widget", "deactivate"}
    args = utils.MakeArg(args, "<widget-id>...", d.WidgetId)
    return args
}

