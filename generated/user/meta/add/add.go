package add

//Adds a meta field.

type Add struct {
    
    User string // <user>
    
    Key string // <key>
    
    Value string // <value>
    
    Format string // [--format=<format>]
    
}

//## OPTIONS
//
//<user>
//: The user login, user email, or user ID of the user to add metadata for.
//
//<key>
//: The metadata key.
//
//<value>
//: The new metadata value.
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
//## EXAMPLES
//
//    # Add user meta
//    $ wp user meta add 123 bio "Mary is an WordPress developer."
//    Success: Added custom field.
//
//