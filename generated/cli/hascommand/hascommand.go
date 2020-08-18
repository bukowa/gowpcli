package hascommand

//Detects if a command exists

type HasCommand struct {
    
    CommandName []string // <command_name>...
    
}

//This commands checks if a command is registered with WP-CLI.
//If the command is found then it returns with exit status 0.
//If the command doesn't exist, then it will exit with status 1.
//
//## OPTIONS
//<command_name>...
//: The command
//
//## EXAMPLES
//
//    # The "site delete" command is registered.
//    $ wp cli has-command "site delete"
//    $ echo $?
//    0
//
//    # The "foo bar" command is not registered.
//    $ wp cli has-command "foo bar"
//    $ echo $?
//    1
//
//