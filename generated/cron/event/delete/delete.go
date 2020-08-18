package delete

//Deletes the next scheduled cron event for the given hook.

type Delete struct {
    
    Hook string // <hook>
    
}

//## OPTIONS
//
//<hook>
//: The hook name.
//
//## EXAMPLES
//
//    # Delete the next scheduled cron event
//    $ wp cron event delete cron_test
//    Success: Deleted 2 instances of the cron event 'cron_test'.
//
//