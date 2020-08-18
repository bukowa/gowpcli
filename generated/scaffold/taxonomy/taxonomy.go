package taxonomy

//Generates PHP code for registering a custom taxonomy.

type Taxonomy struct {
    
    Slug string // <slug>
    
    PostTypes string // [--post_types=<post-types>]
    
    Label string // [--label=<label>]
    
    Textdomain string // [--textdomain=<textdomain>]
    
    Theme bool // [--theme]
    
    Plugin string // [--plugin=<plugin>]
    
    Raw bool // [--raw]
    
    Force bool // [--force]
    
}

//## OPTIONS
//
//<slug>
//: The internal name of the taxonomy.
//
//[--post_types=<post-types>]
//: Post types to register for use with the taxonomy.
//
//[--label=<label>]
//: The text used to translate the update messages.
//
//[--textdomain=<textdomain>]
//: The textdomain to use for the labels.
//
//[--theme]
//: Create a file in the active theme directory, instead of sending to
//STDOUT. Specify a theme with `--theme=<theme>` to have the file placed in that theme.
//
//[--plugin=<plugin>]
//: Create a file in the given plugin's directory, instead of sending to STDOUT.
//
//[--raw]
//: Just generate the `register_taxonomy()` call and nothing else.
//
//[--force]
//: Overwrite files that already exist.
//
//## EXAMPLES
//
//    # Generate PHP code for registering a custom taxonomy and save in a file
//    $ wp scaffold taxonomy venue --post_types=event,presentation > taxonomy.php
//
//