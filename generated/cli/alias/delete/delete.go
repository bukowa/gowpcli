package delete

//Deletes an alias.

type Delete struct {
    
    Key string // <key>
    
    Config string // [--config=<config>]
    
}

//## OPTIONS
//
//<key>
//: Key for the alias.
//
//[--config=<config>]
//: Config file to be considered for operations.
//---
//options:
//  - global
//  - project
//---
//
//## EXAMPLES
//
//    # Delete alias.
//    $ wp cli alias delete @prod
//    Success: Deleted '@prod' alias.
//
//    # Delete project alias.
//    $ wp cli alias delete @prod --config=project
//    Success: Deleted '@prod' alias.
//
//