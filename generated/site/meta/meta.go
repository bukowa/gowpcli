/*
## EXAMPLES
	    # Set site meta
	    $ wp site meta set 123 bio "Mary is a WordPress developer."
	    Success: Updated custom field 'bio'.
	    # Get site meta
	    $ wp site meta get 123 bio
	    Mary is a WordPress developer.
	    # Update site meta
	    $ wp site meta update 123 bio "Mary is an awesome WordPress developer."
	    Success: Updated custom field 'bio'.
	    # Delete site meta
	    $ wp site meta delete 123 bio
	    Success: Deleted custom field.
	
	
 */
package meta


// Meta //Adds, updates, deletes, and lists site custom fields.
type Meta struct {
}

func (m Meta) Args() []string {
    var args = []string{"site", "meta"}
    return args
}

