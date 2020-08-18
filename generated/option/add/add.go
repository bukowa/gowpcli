package add

//Adds a new option value.

type Add struct {
    
    Key string // <key>
    
    Value string // [<value>]
    
    Format string // [--format=<format>]
    
    Autoload string // [--autoload=<autoload>]
    
}

//Errors if the option already exists.
//
//## OPTIONS
//
//<key>
//: The name of the option to add.
//
//[<value>]
//: The value of the option to add. If ommited, the value is read from STDIN.
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
//[--autoload=<autoload>]
//: Should this option be automatically loaded.
//---
//options:
//  - 'yes'
//  - 'no'
//---
//
//## EXAMPLES
//
//    # Create an option by reading a JSON file.
//    $ wp option add my_option --format=json < config.json
//    Success: Added 'my_option' option.
//
//