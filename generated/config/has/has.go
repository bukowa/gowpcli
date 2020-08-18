package has

//Checks whether a specific constant or variable exists in the wp-config.php file.

type Has struct {
    
    Name string // <name>
    
    Type string // [--type=<type>]
    
}

//## OPTIONS
//
//<name>
//: Name of the wp-config.php constant or variable.
//
//[--type=<type>]
//: Type of the config value to set. Defaults to 'all'.
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
//    # Check whether the DB_PASSWORD constant exists in the wp-config.php file.
//    $ wp config has DB_PASSWORD
//
//