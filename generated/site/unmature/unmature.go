package unmature

//Sets one or more sites as immature.

type Unmature struct {
    
    Id []string // <id>...
    
}

//## OPTIONS
//
//<id>...
//: One or more IDs of sites to set as unmature.
//
//## EXAMPLES
//
//    $ wp site general 123
//    Success: Site 123 marked as unmature.
//
//