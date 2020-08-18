/*
## OPTIONS
	[--fields=<fields>]
	: Limit the output to specific object fields.
	[--field=<field>]
	: Prints the value of a single field for each schedule.
	[--format=<format>]
	: Render output in a particular format.
	---
	default: table
	options:
	  - table
	  - csv
	  - ids
	  - json
	  - yaml
	---
## AVAILABLE FIELDS
	These fields will be displayed by default for each cron schedule:
	* name
	* display
	* interval
	There are no additional fields.
## EXAMPLES
	    # List available cron schedules
	    $ wp cron schedule list
	    +------------+-------------+----------+
	    | name       | display     | interval |
	    +------------+-------------+----------+
	    | hourly     | Once Hourly | 3600     |
	    | twicedaily | Twice Daily | 43200    |
	    | daily      | Once Daily  | 86400    |
	    +------------+-------------+----------+
	    # List id of available cron schedule
	    $ wp cron schedule list --fields=name --format=ids
	    hourly twicedaily daily
	
 */
package list
import utils "github.com/bukowa/gowpcli"

// List //List available cron schedules.
type List struct {
    Fields string // [--fields=<fields>]
    Field string // [--field=<field>]
    Format string // [--format=<format>]
}

func (l List) Args() []string {
    var args = []string{"cron", "schedule", "list"}
    args = utils.MakeArg(args, "[--fields=<fields>]", l.Fields)
    args = utils.MakeArg(args, "[--field=<field>]", l.Field)
    args = utils.MakeArg(args, "[--format=<format>]", l.Format)
    return args
}

