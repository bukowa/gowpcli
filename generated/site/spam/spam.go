package spam

//Marks one or more sites as spam.

type Spam struct {
    
    Id []string // <id>...
    
}

//## OPTIONS
//
//<id>...
//: One or more IDs of sites to be marked as spam.
//
//## EXAMPLES
//
//    $ wp site spam 123
//    Success: Site 123 marked as spam.
//
//