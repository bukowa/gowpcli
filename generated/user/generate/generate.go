package generate

//Generates some users.

type Generate struct {
    
    Count string // [--count=<number>]
    
    Role string // [--role=<role>]
    
    Format string // [--format=<format>]
    
}

//Creates a specified number of new users with dummy data.
//
//## OPTIONS
//
//[--count=<number>]
//: How many users to generate?
//---
//default: 100
//---
//
//[--role=<role>]
//: The role of the generated users. Default: default role from WP
//
//[--format=<format>]
//: Render output in a particular format.
//---
//default: progress
//options:
//  - progress
//  - ids
//---
//
//## EXAMPLES
//
//    # Add meta to every generated users.
//    $ wp user generate --format=ids --count=3 | xargs -d ' ' -I % wp user meta add % foo bar
//    Success: Added custom field.
//    Success: Added custom field.
//    Success: Added custom field.
//
//