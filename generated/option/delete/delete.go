package delete

//Deletes an option.

type Delete struct {
    
    Key []string // <key>...
    
}

//## OPTIONS
//
//<key>...
//: Key for the option.
//
//## EXAMPLES
//
//    # Delete an option.
//    $ wp option delete my_option
//    Success: Deleted 'my_option' option.
//
//    # Delete multiple options.
//    $ wp option delete option_one option_two option_three
//    Success: Deleted 'option_one' option.
//    Success: Deleted 'option_two' option.
//    Warning: Could not delete 'option_three' option. Does it exist?
//
//