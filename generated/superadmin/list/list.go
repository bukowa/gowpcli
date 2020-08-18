package list

//Lists users with super admin capabilities.

type List struct {
    
    Format string // [--format=<format>]
    
}

//## OPTIONS
//
//[--format=<format>]
//: Render output in a particular format.
//---
//default: list
//options:
//  - list
//  - table
//  - csv
//  - json
//  - count
//  - yaml
//---
//
//## EXAMPLES
//
//    # List user with super-admin capabilities
//    $ wp super-admin list
//    supervisor
//    administrator
//
//