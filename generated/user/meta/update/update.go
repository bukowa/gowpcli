package update

//Updates a meta field.

type Update struct {
    
    User string // <user>
    
    Key string // <key>
    
    Value string // <value>
    
    Format string // [--format=<format>]
    
}

//## OPTIONS
//
//<user>
//: The user login, user email, or user ID of the user to update metadata for.
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
//    # Update user meta
//    $ wp user meta update 123 bio "Mary is an awesome WordPress developer."
//    Success: Updated custom field 'bio'.
//
//