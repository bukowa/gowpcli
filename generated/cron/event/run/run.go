package run

//Runs the next scheduled cron event for the given hook.

type Run struct {
    
    Hook []string // [<hook>...]
    
    DueNow bool // [--due-now]
    
    All bool // [--all]
    
}

//## OPTIONS
//
//[<hook>...]
//: One or more hooks to run.
//
//[--due-now]
//: Run all hooks due right now.
//
//[--all]
//: Run all hooks.
//
//## EXAMPLES
//
//    # Run all cron events due right now
//    $ wp cron event run --due-now
//    Success: Executed a total of 2 cron events.
//
//