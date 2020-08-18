/*
## EXAMPLES
	    # Test WP Cron spawning system
	    $ wp cron test
	    Success: WP-Cron spawning is working as expected.
	
 */
package cron


// Cron //Tests, runs, and deletes WP-Cron events; manages WP-Cron schedules.
type Cron struct {
}

func (c Cron) Args() []string {
    var args = []string{"cron"}
    return args
}

