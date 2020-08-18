/*
## INFO
	See the WordPress [Rewrite API](https://codex.wordpress.org/Rewrite_API) and
	[WP Rewrite](https://codex.wordpress.org/Class_Reference/WP_Rewrite) class reference.
## EXAMPLES
	    # Flush rewrite rules
	    $ wp rewrite flush
	    Success: Rewrite rules flushed.
	    # Update permalink structure
	    $ wp rewrite structure '/%year%/%monthnum%/%postname%'
	    Success: Rewrite structure set.
	    # List rewrite rules
	    $ wp rewrite list --format=csv
	    match,query,source
	    ^wp-json/?$,index.php?rest_route=/,other
	    ^wp-json/(.*)?,index.php?rest_route=/$matches[1],other
	    category/(.+?)/feed/(feed|rdf|rss|rss2|atom)/?$,index.php?category_name=$matches[1]&feed=$matches[2],category
	    category/(.+?)/(feed|rdf|rss|rss2|atom)/?$,index.php?category_name=$matches[1]&feed=$matches[2],category
	    category/(.+?)/embed/?$,index.php?category_name=$matches[1]&embed=true,category
	
 */
package rewrite


// Rewrite //Lists or flushes the site's rewrite rules, updates the permalink structure.
type Rewrite struct {
}

func (r Rewrite) Args() []string {
    var args = []string{"rewrite"}
    return args
}

