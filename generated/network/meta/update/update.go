package update

//Update a meta field.

type Update struct {
    
    Id string // <id>
    
    Key string // <key>
    
    Value string // [<value>]
    
    Format string // [--format=<format>]
    
}

//## OPTIONS
//
//<id>
//: The ID of the object.
//
//<key>
//: The name of the meta field to update.
//
//[<value>]
//: The new value. If omitted, the value is read from STDIN.
//
//[--format=<format>]
//: The serialization format for the value.
//---
//default: plaintext
//options:
//  - plaintext
//  - json
//---
//
//