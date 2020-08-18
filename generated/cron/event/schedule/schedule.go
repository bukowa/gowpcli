/*
## OPTIONS
	<hook>
	: The hook name.
	[<next-run>]
	: A Unix timestamp or an English textual datetime description compatible with `strtotime()`. Defaults to now.
	[<recurrence>]
	: How often the event should recur. See `wp cron schedule list` for available schedule names. Defaults to no recurrence.
	[--<field>=<value>]
	: Associative args for the event.
## EXAMPLES
	    # Schedule a new cron event
	    $ wp cron event schedule cron_test
	    Success: Scheduled event with hook 'cron_test' for 2016-05-31 10:19:16 GMT.
	    # Schedule new cron event with hourly recurrence
	    $ wp cron event schedule cron_test now hourly
	    Success: Scheduled event with hook 'cron_test' for 2016-05-31 10:20:32 GMT.
	    # Schedule new cron event and pass associative arguments
	    $ wp cron event schedule cron_test '+1 hour' --foo=1 --bar=2
	    Success: Scheduled event with hook 'cron_test' for 2016-05-31 11:21:35 GMT.
	
 */
package schedule
import utils "github.com/bukowa/gowpcli"

// Schedule //Schedules a new cron event.
type Schedule struct {
    Hook string // <hook>
    NextRun string // [<next-run>]
    Recurrence string // [<recurrence>]
    FieldMap map[string]string // [--<field>=<value>]
}

func (s Schedule) Args() []string {
    var args = []string{"cron", "event", "schedule"}
    args = utils.MakeArg(args, "<hook>", s.Hook)
    args = utils.MakeArg(args, "[<next-run>]", s.NextRun)
    args = utils.MakeArg(args, "[<recurrence>]", s.Recurrence)
    args = utils.MakeArg(args, "[--<field>=<value>]", s.FieldMap)
    return args
}

