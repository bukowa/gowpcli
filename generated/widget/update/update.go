/*
## OPTIONS
	<widget-id>
	: Unique ID for the widget
	[--<field>=<value>]
	: Field to update, with its new value
## EXAMPLES
	    # Change calendar-1 widget title to "Our Calendar"
	    $ wp widget update calendar-1 --title="Our Calendar"
	    Success: Widget updated.
	
 */
package update
import utils "github.com/bukowa/gowpcli"

// Update //Updates options for an existing widget.
type Update struct {
    WidgetId string // <widget-id>
    FieldMap map[string]string // [--<field>=<value>]
}

func (u Update) Args() []string {
    var args = []string{"widget", "update"}
    args = utils.MakeArg(args, "<widget-id>", u.WidgetId)
    args = utils.MakeArg(args, "[--<field>=<value>]", u.FieldMap)
    return args
}

