package delete

//Deletes a specific constant or variable from the wp-config.php file.

type Delete struct {
    
    Name string // <name>
    
    Type string // [--type=<type>]
    
}

//## OPTIONS
//
//<name>
//: Name of the wp-config.php constant or variable.
//
//[--type=<type>]
//: Type of the config value to delete. Defaults to 'all'.
//---
//default: all
//options:
//  - constant
//  - variable
//  - all
//---
//
//## EXAMPLES
//
//    # Delete the COOKIE_DOMAIN constant from the wp-config.php file.
//    $ wp config delete COOKIE_DOMAIN
//
//