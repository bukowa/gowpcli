package add

//Adds a widget to a sidebar.

type Add struct {
    
    Name string // <name>
    
    SidebarId string // <sidebar-id>
    
    Position string // [<position>]
    
    FieldMap map[string]string // [--<field>=<value>]
    
}

//Creates a new widget entry in the database, and associates it with the
//sidebar.
//
//## OPTIONS
//
//<name>
//: Widget name.
//
//<sidebar-id>
//: ID for the corresponding sidebar.
//
//[<position>]
//: Widget's current position within the sidebar. Defaults to last
//
//[--<field>=<value>]
//: Widget option to add, with its new value
//
//## EXAMPLES
//
//    # Add a new calendar widget to sidebar-1 with title "Calendar"
//    $ wp widget add calendar sidebar-1 2 --title="Calendar"
//    Success: Added widget to sidebar.
//
//