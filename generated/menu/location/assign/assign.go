package assign

//Assigns a location to a menu.

type Assign struct {
    
    Menu string // <menu>
    
    Location string // <location>
    
}

//## OPTIONS
//
//<menu>
//: The name, slug, or term ID for the menu.
//
//<location>
//: Location's slug.
//
//## EXAMPLES
//
//    $ wp menu location assign primary-menu primary
//    Success: Assigned location primary to menu primary-menu.
//
//