/*
## INFO
	
	
 */
package cache


// Cache //Finds, triggers, and deletes oEmbed caches.
type Cache struct {
}

func (c Cache) Args() []string {
    var args = []string{"embed", "cache"}
    return args
}

