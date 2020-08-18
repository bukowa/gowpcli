package create

//Creates a new menu.

type Create struct {
    
    MenuName string // <menu-name>
    
    Porcelain bool // [--porcelain]
    
}

//## OPTIONS
//
//<menu-name>
//: A descriptive name for the menu.
//
//[--porcelain]
//: Output just the new menu id.
//
//## EXAMPLES
//
//    $ wp menu create "My Menu"
//    Success: Created menu 200.
//
//