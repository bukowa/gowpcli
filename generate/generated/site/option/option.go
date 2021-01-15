/*
## EXAMPLES
	    # Get site registration
	    $ wp site option get registration
	    none
	    # Add site option
	    $ wp site option add my_option foobar
	    Success: Added 'my_option' site option.
	    # Update site option
	    $ wp site option update my_option '{"foo": "bar"}' --format=json
	    Success: Updated 'my_option' site option.
	    # Delete site option
	    $ wp site option delete my_option
	    Success: Deleted 'my_option' site option.
	
	
 */
package option


// Option //Adds, updates, deletes, and lists site options in a multisite installation.
type Option struct {
}

func (o Option) Args() []string {
    var args = []string{"site", "option"}
    return args
}

