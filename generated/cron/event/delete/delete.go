/*
## OPTIONS
	<hook>
	: The hook name.
## EXAMPLES
	    # Delete the next scheduled cron event
	    $ wp cron event delete cron_test
	    Success: Deleted 2 instances of the cron event 'cron_test'.
	
 */
package delete
import utils "github.com/bukowa/gowpcli"

// Delete //Deletes the next scheduled cron event for the given hook.
type Delete struct {
    Hook string // <hook>
}

func (d Delete) Args() []string {
    var args = []string{"cron", "event", "delete"}
    args = utils.MakeArg(args, "<hook>", d.Hook)
    return args
}

