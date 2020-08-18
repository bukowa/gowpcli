package update

//Updates an alias.

type Update struct {
    
    Key string // <key>
    
    SetUser string // [--set-user=<user>]
    
    SetUrl string // [--set-url=<url>]
    
    SetPath string // [--set-path=<path>]
    
    SetSsh string // [--set-ssh=<ssh>]
    
    SetHttp string // [--set-http=<http>]
    
    Grouping string // [--grouping=<grouping>]
    
    Config string // [--config=<config>]
    
}

//## OPTIONS
//
//<key>
//: Key for the alias.
//
//[--set-user=<user>]
//: Set user for alias.
//
//[--set-url=<url>]
//: Set url for alias.
//
//[--set-path=<path>]
//: Set path for alias.
//
//[--set-ssh=<ssh>]
//: Set ssh for alias.
//
//[--set-http=<http>]
//: Set http for alias.
//
//[--grouping=<grouping>]
//: For grouping multiple aliases.
//
//[--config=<config>]
//: Config file to be considered for operations.
//---
//options:
//  - global
//  - project
//---
//
//## EXAMPLES
//
//    # Update alias.
//    $ wp cli alias update @prod --set-user=newuser --set-path=/new/path/to/wordpress/install/
//    Success: Updated 'prod' alias.
//
//    # Update project alias.
//    $ wp cli alias update @prod --set-user=newuser --set-path=/new/path/to/wordpress/install/ --config=project
//    Success: Updated 'prod' alias.
//
//