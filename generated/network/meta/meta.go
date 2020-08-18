/*
## EXAMPLES
	    # Get a list of super-admins
	    $ wp network meta get 1 site_admins
	    array (
	      0 => 'supervisor',
	    )
	
 */
package meta


// Meta //Gets, adds, updates, deletes, and lists network custom fields.
type Meta struct {
}

func (m Meta) Args() []string {
    var args = []string{"network", "meta"}
    return args
}

