package move

//Moves the position of a widget.

type Move struct {
    
    WidgetId string // <widget-id>
    
    Position string // [--position=<position>]
    
    SidebarId string // [--sidebar-id=<sidebar-id>]
    
}

//Changes the order of a widget in its existing sidebar, or moves it to a
//new sidebar.
//
//## OPTIONS
//
//<widget-id>
//: Unique ID for the widget
//
//[--position=<position>]
//: Assign the widget to a new position.
//
//[--sidebar-id=<sidebar-id>]
//: Assign the widget to a new sidebar
//
//## EXAMPLES
//
//    # Change position of widget
//    $ wp widget move recent-comments-2 --position=2
//    Success: Widget moved.
//
//    # Move widget to Inactive Widgets
//    $ wp widget move recent-comments-2 --sidebar-id=wp_inactive_widgets
//    Success: Widget moved.
//
//