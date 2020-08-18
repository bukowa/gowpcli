package get

//Gets meta field value.

type Get struct {
    
    User string // <user>
    
    Key string // <key>
    
    Format string // [--format=<format>]
    
}

//## OPTIONS
//
//<user>
//: The user login, user email, or user ID of the user to get metadata for.
//
//<key>
//: The metadata key.
//
//[--format=<format>]
//: Render output in a particular format.
//---
//default: table
//options:
//  - table
//  - csv
//  - json
//  - yaml
//---
//
//## EXAMPLES
//
//    # Get user meta
//    $ wp user meta get 123 bio
//    Mary is an WordPress developer.
//
//