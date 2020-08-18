package help

//Gets help on WP-CLI, or on a specific command.

type Help struct {
    
    Command []string // [<command>...]
    
}

//## OPTIONS
//
//[<command>...]
//: Get help on a specific command.
//
//## EXAMPLES
//
//    # get help for `core` command
//    wp help core
//
//    # get help for `core download` subcommand
//    wp help core download
//
//