/*
## OPTIONS
	<hook>
	: Name of the hook for which all events should be unscheduled.
## EXAMPLES
	    # Unschedule a cron event on given hook.
	    $ wp cron event unschedule cron_test
	    Success: Unscheduled 2 events with hook 'cron_test'.
	
 */
package unschedule
import utils "github.com/bukowa/gowpcli"

// Unschedule //Unschedules all cron events for a given hook.
type Unschedule struct {
    Hook string // <hook>
}

func (u Unschedule) Args() []string {
    var args = []string{"cron", "event", "unschedule"}
    args = utils.MakeArg(args, "<hook>", u.Hook)
    return args
}

