/*
## EXAMPLES
	    # Set post meta
	    $ wp post meta set 123 _wp_page_template about.php
	    Success: Updated custom field '_wp_page_template'.
	    # Get post meta
	    $ wp post meta get 123 _wp_page_template
	    about.php
	    # Update post meta
	    $ wp post meta update 123 _wp_page_template contact.php
	    Success: Updated custom field '_wp_page_template'.
	    # Delete post meta
	    $ wp post meta delete 123 _wp_page_template
	    Success: Deleted custom field.
	
	
 */
package meta


// Meta //Adds, updates, deletes, and lists post custom fields.
type Meta struct {
}

func (m Meta) Args() []string {
    var args = []string{"post", "meta"}
    return args
}

