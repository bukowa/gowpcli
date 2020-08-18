package list

//Lists capabilities for a given role.

type List struct {
    
    Role string // <role>
    
    Format string // [--format=<format>]
    
    ShowGrant bool // [--show-grant]
    
}

//## OPTIONS
//
//<role>
//: Key for the role.
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
//[--show-grant]
//: Display all capabilities defined for a role including grant.
//---
//default: false
//---
//
//## EXAMPLES
//
//    # Display alphabetical list of Contributor capabilities.
//    $ wp cap list 'contributor' | sort
//    delete_posts
//    edit_posts
//    level_0
//    level_1
//    read
//
//