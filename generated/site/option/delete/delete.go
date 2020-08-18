package delete

//Deletes a site option.

type Delete struct {
    
    Key string // <key>
    
}

//## OPTIONS
//
//<key>
//: Key for the site option.
//
//## EXAMPLES
//
//    $ wp site option delete my_option
//    Success: Deleted 'my_option' site option.
//
//