package unarchive

//Unarchives one or more sites.

type Unarchive struct {
    
    Id []string // <id>...
    
}

//## OPTIONS
//
//<id>...
//: One or more IDs of sites to unarchive.
//
//## EXAMPLES
//
//    $ wp site unarchive 123
//    Success: Site 123 unarchived.
//
//