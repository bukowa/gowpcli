package private

//Sets one or more sites as private.

type Private struct {
    
    Id []string // <id>...
    
}

//## OPTIONS
//
//<id>...
//: One or more IDs of sites to set as private.
//
//## EXAMPLES
//
//    $ wp site private 123
//    Success: Site 123 marked as private.
//
//