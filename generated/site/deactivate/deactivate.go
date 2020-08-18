package deactivate

//Deactivates one or more sites.

type Deactivate struct {
    
    Id []string // <id>...
    
}

//## OPTIONS
//
//<id>...
//: One or more IDs of sites to deactivate.
//
//## EXAMPLES
//
//    $ wp site deactivate 123
//    Success: Site 123 deactivated.
//
//