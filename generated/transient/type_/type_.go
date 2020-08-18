/*
## INFO
	Indicates whether the transients API is using an object cache or the
	database.
	For a more complete explanation of the transient cache, including the
	network|site cache, please see docs for `wp transient`.
## EXAMPLES
	    $ wp transient type
	    Transients are saved to the database.
	
 */
package type_


// Type_ //Determines the type of transients implementation.
type Type_ struct {
}

func (t Type_) Args() []string {
    var args = []string{"transient", "type"}
    return args
}

