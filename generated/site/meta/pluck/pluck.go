package pluck

//Get a nested value from a meta field.

type Pluck struct {
    
    Id string // <id>
    
    Key string // <key>
    
    KeyPath []string // <key-path>...
    
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
//<key-path>...
//: The name(s) of the keys within the value to locate the value to pluck.
//
//[--format=<format>]
//: The output format of the value.
//---
//default: plaintext
//options:
//  - plaintext
//  - json
//  - yaml
//
//