package get

//Gets the value for an option.

type Get struct {
    
    Key string // <key>
    
    Format string // [--format=<format>]
    
}

//## OPTIONS
//
//<key>
//: Key for the option.
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
//    # Get option.
//    $ wp option get home
//    http://example.com
//
//    # Get blog description.
//    $ wp option get blogdescription
//    A random blog description
//
//    # Get blog name
//    $ wp option get blogname
//    A random blog name
//
//    # Get admin email.
//    $ wp option get admin_email
//    someone@example.com
//
//    # Get option in JSON format.
//    $ wp option get active_plugins --format=json
//    {"0":"dynamically-dynamic-sidebar\/dynamically-dynamic-sidebar.php","1":"monster-widget\/monster-widget.php","2":"show-current-template\/show-current-template.php","3":"theme-check\/theme-check.php","5":"wordpress-importer\/wordpress-importer.php"}
//
//