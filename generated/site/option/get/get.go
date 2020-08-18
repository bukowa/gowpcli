package get

//Gets a site option.

type Get struct {
    
    Key string // <key>
    
    Format string // [--format=<format>]
    
}

//## OPTIONS
//
//<key>
//: Key for the site option.
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
//## EXAMPLES
//
//    # Get site upload filetypes
//    $ wp site option get upload_filetypes
//    jpg jpeg png gif mov avi mpg
//
//