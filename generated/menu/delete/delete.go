package delete

//Deletes one or more menus.

type Delete struct {
    
    Menu []string // <menu>...
    
}

//## OPTIONS
//
//<menu>...
//: The name, slug, or term ID for the menu(s).
//
//## EXAMPLES
//
//    $ wp menu delete "My Menu"
//    Success: 1 menu deleted.
//
//