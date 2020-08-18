/*
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
	
	
 */
package schedule


// Schedule //Gets WP-Cron schedules.
type Schedule struct {
}

func (s Schedule) Args() []string {
    var args = []string{"cron", "schedule"}
    return args
}

