/*
## EXAMPLES
	    # Launch system editor to edit wp-config.php file
	    $ wp config edit
	    # Edit wp-config.php file in a specific editor
	    $ EDITOR=vim wp config edit
	
 */
package edit


// Edit //Launches system editor to edit the wp-config.php file.
type Edit struct {
}

func (e Edit) Args() []string {
    var args = []string{"config", "edit"}
    return args
}

