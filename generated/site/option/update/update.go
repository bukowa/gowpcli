package update

//Updates a site option.

type Update struct {
    
    Key string // <key>
    
    Value string // [<value>]
    
    Format string // [--format=<format>]
    
}

//## OPTIONS
//
//<key>
//: The name of the site option to update.
//
//[<value>]
//: The new value. If ommited, the value is read from STDIN.
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
//    # Update a site option by reading from a file
//    $ wp site option update my_option < value.txt
//    Success: Updated 'my_option' site option.
//
//