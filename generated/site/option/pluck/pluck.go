package pluck

//Gets a nested value from an option.

type Pluck struct {
    
    Key string // <key>
    
    KeyPath []string // <key-path>...
    
    Format string // [--format=<format>]
    
}

//## OPTIONS
//
//<key>
//: The option name.
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