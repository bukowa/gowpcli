package set

//Sets the value of a specific constant or variable defined in wp-config.php file.

type Set struct {
    
    Name string // <name>
    
    Value string // <value>
    
    Add bool // [--add]
    
    Raw bool // [--raw]
    
    Anchor string // [--anchor=<anchor>]
    
    Placement string // [--placement=<placement>]
    
    Separator string // [--separator=<separator>]
    
    Type string // [--type=<type>]
    
}

//## OPTIONS
//
//<name>
//: Name of the wp-config.php constant or variable.
//
//<value>
//: Value to set the wp-config.php constant or variable to.
//
//[--add]
//: Add the value if it doesn't exist yet.
//This is the default behavior, override with --no-add.
//
//[--raw]
//: Place the value into the wp-config.php file as is, instead of as a quoted string.
//
//[--anchor=<anchor>]
//: Anchor string where additions of new values are anchored around.
//Defaults to "/* That's all, stop editing!".
//
//[--placement=<placement>]
//: Where to place the new values in relation to the anchor string.
//---
//default: 'before'
//options:
//  - before
//  - after
//---
//
//[--separator=<separator>]
//: Separator string to put between an added value and its anchor string.
//The following escape sequences will be recognized and properly interpreted: '\n' => newline, '\r' => carriage return, '\t' => tab.
//Defaults to a single EOL ("\n" on *nix and "\r\n" on Windows).
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
//    # Set the WP_DEBUG constant to true.
//    $ wp config set WP_DEBUG true --raw
//
//