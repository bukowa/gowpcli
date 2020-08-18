/*
## OPTIONS
## EXAMPLES
	    # Get new salts for your wp-config.php file
	    $ wp config shuffle-salts
	    Success: Shuffled the salt keys.
	
 */
package shufflesalts


// ShuffleSalts //Refreshes the salts defined in the wp-config.php file.
type ShuffleSalts struct {
}

func (s ShuffleSalts) Args() []string {
    var args = []string{"config", "shuffle-salts"}
    return args
}

