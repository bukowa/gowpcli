package cache

//Manages the internal WP-CLI cache,.

type Cache struct {
    
}

//## EXAMPLES
//
//    # Remove all cached files.
//    $ wp cli cache clear
//    Success: Cache cleared.
//
//    # Remove all cached files except for the newest version of each one.
//    $ wp cli cache prune
//    Success: Cache pruned.
//
//
//
//