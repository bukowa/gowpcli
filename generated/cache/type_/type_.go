/*
## INFO
	Note that the guesses made by this function are based on the
	WP_Object_Cache classes that define the 3rd party object cache extension.
	Changes to those classes could render problems with this function's
	ability to determine which object cache is being used.
## EXAMPLES
	    # Check cache type.
	    $ wp cache type
	    Default
	
 */
package type_


// Type_ //Attempts to determine which object cache is being used.
type Type_ struct {
}

func (t Type_) Args() []string {
    var args = []string{"cache", "type"}
    return args
}

