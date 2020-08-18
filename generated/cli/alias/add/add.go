package add

//Creates an alias.

type Add struct {
    
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
//default: global
//options:
//  - global
//  - project
//---
//
//## EXAMPLES
//
//    # Add alias to global config.
//    $ wp cli alias add @prod  --set-ssh=login@host --set-path=/path/to/wordpress/install/ --set-user=wpcli
//    Success: Added '@prod' alias.
//
//    # Add alias to project config.
//    $ wp cli alias add @prod --set-ssh=login@host --set-path=/path/to/wordpress/install/ --set-user=wpcli --config=project
//    Success: Added '@prod' alias.
//
//    # Add group of aliases.
//    $ wp cli alias add @multiservers --grouping=servera,serverb
//    Success: Added '@multiservers' alias.
//
//