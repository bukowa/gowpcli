/*
## EXAMPLES
	    $ wp cli cache clear
	    Success: Cache cleared.
	
 */
package clear


// Clear //Clears the internal cache.
type Clear struct {
}

func (c Clear) Args() []string {
    var args = []string{"cli", "cache", "clear"}
    return args
}

