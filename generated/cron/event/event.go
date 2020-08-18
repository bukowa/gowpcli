/*
## EXAMPLES
	    # Schedule a new cron event
	    $ wp cron event schedule cron_test
	    Success: Scheduled event with hook 'cron_test' for 2016-05-31 10:19:16 GMT.
	    # Run all cron events due right now
	    $ wp cron event run --due-now
	    Success: Executed a total of 2 cron events.
	    # Delete the next scheduled cron event
	    $ wp cron event delete cron_test
	    Success: Deleted 2 instances of the cron event 'cron_test'.
	    # List scheduled cron events in JSON
	    $ wp cron event list --fields=hook,next_run --format=json
	    [{"hook":"wp_version_check","next_run":"2016-05-31 10:15:13"},{"hook":"wp_update_plugins","next_run":"2016-05-31 10:15:13"},{"hook":"wp_update_themes","next_run":"2016-05-31 10:15:14"}]
	
	
 */
package event


// Event //Schedules, runs, and deletes WP-Cron events.
type Event struct {
}

func (e Event) Args() []string {
    var args = []string{"cron", "event"}
    return args
}

