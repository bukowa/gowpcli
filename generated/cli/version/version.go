/*
## EXAMPLES
	    # Display CLI version.
	    $ wp cli version
	    WP-CLI 0.24.1
	
 */
package version


// Version //Prints WP-CLI version.
type Version struct {
}

func (v Version) Args() []string {
    var args = []string{"cli", "version"}
    return args
}

