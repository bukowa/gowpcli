package patch

//Updates a nested value in an option.

type Patch struct {
    
    Action string // <action>
    
    Key string // <key>
    
    KeyPath []string // <key-path>...
    
    Value string // [<value>]
    
    Format string // [--format=<format>]
    
}

//## OPTIONS
//
//<action>
//: Patch action to perform.
//---
//options:
//  - insert
//  - update
//  - delete
//---
//
//<key>
//: The option name.
//
//<key-path>...
//: The name(s) of the keys within the value to locate the value to patch.
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