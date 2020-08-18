package delete

//Deletes one or more widgets from a sidebar.

type Delete struct {
    
    WidgetId []string // <widget-id>...
    
}

//## OPTIONS
//
//<widget-id>...
//: Unique ID for the widget(s)
//
//## EXAMPLES
//
//    # Delete the recent-comments-2 widget from its sidebar.
//    $ wp widget delete recent-comments-2
//    Success: Deleted 1 of 1 widgets.
//
//