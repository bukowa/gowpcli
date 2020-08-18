package get

//Get meta field value.

type Get struct {
    
    Id string // <id>
    
    Key string // <key>
    
    Format string // [--format=<format>]
    
}

//## OPTIONS
//
//<id>
//: The ID of the object.
//
//<key>
//: The name of the meta field to get.
//
//[--format=<format>]
//: Get value in a particular format.
//---
//default: var_export
//options:
//  - var_export
//  - json
//  - yaml
//---
//
//