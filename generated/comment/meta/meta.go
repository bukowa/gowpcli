/*
## EXAMPLES
	    # Set comment meta
	    $ wp comment meta set 123 description "Mary is a WordPress developer."
	    Success: Updated custom field 'description'.
	    # Get comment meta
	    $ wp comment meta get 123 description
	    Mary is a WordPress developer.
	    # Update comment meta
	    $ wp comment meta update 123 description "Mary is an awesome WordPress developer."
	    Success: Updated custom field 'description'.
	    # Delete comment meta
	    $ wp comment meta delete 123 description
	    Success: Deleted custom field.
	
	
 */
package meta


// Meta //Adds, updates, deletes, and lists comment custom fields.
type Meta struct {
}

func (m Meta) Args() []string {
    var args = []string{"comment", "meta"}
    return args
}

