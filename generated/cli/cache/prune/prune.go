package prune

//Prunes the internal cache.

type Prune struct {
    
}

//Removes all cached files except for the newest version of each one.
//
//## EXAMPLES
//
//    $ wp cli cache prune
//    Success: Cache pruned.
//
//