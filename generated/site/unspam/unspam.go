package unspam

//Removes one or more sites from spam.

type Unspam struct {
    
    Id []string // <id>...
    
}

//## OPTIONS
//
//<id>...
//: One or more IDs of sites to remove from spam.
//
//## EXAMPLES
//
//    $ wp site unspam 123
//    Success: Site 123 removed from spam.
//
//