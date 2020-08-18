/*
## EXAMPLES
	    $ wp maintenance-mode deactivate
	    Disabling Maintenance mode...
	    Success: Deactivated Maintenance mode.
	
 */
package deactivate


// Deactivate //Deactivates maintenance mode.
type Deactivate struct {
}

func (d Deactivate) Args() []string {
    var args = []string{"maintenance-mode", "deactivate"}
    return args
}

