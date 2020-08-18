/*
## INFO
	By default, the WP Object Cache exists in PHP memory for the length of the
	request (and is emptied at the end). Use a persistent object cache drop-in
	to persist the object cache between requests.
	[Read the codex article](https://codex.wordpress.org/Class_Reference/WP_Object_Cache)
	for more detail.
## EXAMPLES
	    # Set cache.
	    $ wp cache set my_key my_value my_group 300
	    Success: Set object 'my_key' in group 'my_group'.
	    # Get cache.
	    $ wp cache get my_key my_group
	    my_value
	
 */
package cache


// Cache //Adds, removes, fetches, and flushes the WP Object Cache object.
type Cache struct {
}

func (c Cache) Args() []string {
    var args = []string{"cache"}
    return args
}

