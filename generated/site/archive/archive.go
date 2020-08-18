package archive

//Archives one or more sites.

type Archive struct {
    
    Id []string // <id>...
    
}

//## OPTIONS
//
//<id>...
//: One or more IDs of sites to archive.
//
//## EXAMPLES
//
//    $ wp site archive 123
//    Success: Site 123 archived.
//
//