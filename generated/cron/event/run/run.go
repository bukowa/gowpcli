/*
## OPTIONS
	[<hook>...]
	: One or more hooks to run.
	[--due-now]
	: Run all hooks due right now.
	[--all]
	: Run all hooks.
## EXAMPLES
	    # Run all cron events due right now
	    $ wp cron event run --due-now
	    Success: Executed a total of 2 cron events.
	
 */
package run
import utils "github.com/bukowa/gowpcli"

// Run //Runs the next scheduled cron event for the given hook.
type Run struct {
    Hook []string // [<hook>...]
    DueNow bool // [--due-now]
    All bool // [--all]
}

func (r Run) Args() []string {
    var args = []string{"cron", "event", "run"}
    args = utils.MakeArg(args, "[<hook>...]", r.Hook)
    args = utils.MakeArg(args, "[--due-now]", r.DueNow)
    args = utils.MakeArg(args, "[--all]", r.All)
    return args
}

