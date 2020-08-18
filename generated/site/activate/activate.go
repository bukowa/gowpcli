package activate

//Activates one or more sites.

type Activate struct {
    
    Id []string // <id>...
    
}

//## OPTIONS
//
//<id>...
//: One or more IDs of sites to activate.
//
//## EXAMPLES
//
//    $ wp site activate 123
//    Success: Site 123 activated.
//
//