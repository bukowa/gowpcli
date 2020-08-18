/*
## EXAMPLES
	    $ wp maintenance-mode status
	    Maintenance mode is active.
	
 */
package status


// Status //Displays maintenance mode status.
type Status struct {
}

func (s Status) Args() []string {
    var args = []string{"maintenance-mode", "status"}
    return args
}

