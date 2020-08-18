/*
## EXAMPLES
	    # Get a list of super-admins
	    $ wp network meta get 1 site_admins
	    array (
	      0 => 'supervisor',
	    )
	
 */
package network


// Network //Perform network-wide operations.
type Network struct {
}

func (n Network) Args() []string {
    var args = []string{"network"}
    return args
}

