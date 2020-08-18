/*
## EXAMPLES
	    # Set term meta
	    $ wp term meta set 123 bio "Mary is a WordPress developer."
	    Success: Updated custom field 'bio'.
	    # Get term meta
	    $ wp term meta get 123 bio
	    Mary is a WordPress developer.
	    # Update term meta
	    $ wp term meta update 123 bio "Mary is an awesome WordPress developer."
	    Success: Updated custom field 'bio'.
	    # Delete term meta
	    $ wp term meta delete 123 bio
	    Success: Deleted custom field.
	
	
 */
package meta


// Meta //Adds, updates, deletes, and lists term custom fields.
type Meta struct {
}

func (m Meta) Args() []string {
    var args = []string{"term", "meta"}
    return args
}

