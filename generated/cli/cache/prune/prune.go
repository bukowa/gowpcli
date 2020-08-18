/*
## INFO
	Removes all cached files except for the newest version of each one.
## EXAMPLES
	    $ wp cli cache prune
	    Success: Cache pruned.
	
 */
package prune


// Prune //Prunes the internal cache.
type Prune struct {
}

func (p Prune) Args() []string {
    var args = []string{"cli", "cache", "prune"}
    return args
}

