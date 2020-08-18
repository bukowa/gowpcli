package add

//Add a meta field.

type Add struct {
    
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
//: The name of the meta field to create.
//
//[<value>]
//: The value of the meta field. If omitted, the value is read from STDIN.
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