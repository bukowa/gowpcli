package remove

//Removes a location from a menu.

type Remove struct {
    
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
//    $ wp menu location remove primary-menu primary
//    Success: Removed location from menu.
//
//