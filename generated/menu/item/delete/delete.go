package delete

//Deletes one or more items from a menu.

type Delete struct {
    
    DbId []string // <db-id>...
    
}

//## OPTIONS
//
//<db-id>...
//: Database ID for the menu item(s).
//
//## EXAMPLES
//
//    $ wp menu item delete 45
//    Success: 1 menu item deleted.
//
//