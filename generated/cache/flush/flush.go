/*
## INFO
	For WordPress multisite instances using a persistent object cache,
	flushing the object cache will typically flush the cache for all sites.
	Beware of the performance impact when flushing the object cache in
	production.
	Errors if the object cache can't be flushed.
## EXAMPLES
	    # Flush cache.
	    $ wp cache flush
	    Success: The cache was flushed.
	
 */
package flush


// Flush //Flushes the object cache.
type Flush struct {
}

func (f Flush) Args() []string {
    var args = []string{"cache", "flush"}
    return args
}

