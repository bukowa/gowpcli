package mature

//Sets one or more sites as mature.

type Mature struct {
    
    Id []string // <id>...
    
}

//## OPTIONS
//
//<id>...
//: One or more IDs of sites to set as mature.
//
//## EXAMPLES
//
//    $ wp site mature 123
//    Success: Site 123 marked as mature.
//
//