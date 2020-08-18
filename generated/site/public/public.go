package public

//Sets one or more sites as public.

type Public struct {
    
    Id []string // <id>...
    
}

//## OPTIONS
//
//<id>...
//: One or more IDs of sites to set as public.
//
//## EXAMPLES
//
//    $ wp site public 123
//    Success: Site 123 marked as public.
//
//