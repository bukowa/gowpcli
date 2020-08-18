/*
## EXAMPLES
	    $ wp maintenance-mode is-active
	    $ echo $?
	    1
	
 */
package isactive


// IsActive //Detects maintenance mode status.
type IsActive struct {
}

func (i IsActive) Args() []string {
    var args = []string{"maintenance-mode", "is-active"}
    return args
}

