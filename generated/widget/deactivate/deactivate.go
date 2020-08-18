package deactivate

//Deactivates one or more widgets from an active sidebar.

type Deactivate struct {
    
    WidgetId []string // <widget-id>...
    
}

//Moves widgets to Inactive Widgets.
//
//## OPTIONS
//
//<widget-id>...
//: Unique ID for the widget(s)
//
//## EXAMPLES
//
//    # Deactivate the recent-comments-2 widget.
//    $ wp widget deactivate recent-comments-2
//    Success: 1 widget deactivated.
//
//