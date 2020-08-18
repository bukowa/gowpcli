/*
## INFO
	Display the database table prefix, as defined by the database handler's interpretation of the current site.
## EXAMPLES
	    $ wp db prefix
	    wp_
	
 */
package prefix


// Prefix //Displays the database table prefix.
type Prefix struct {
}

func (p Prefix) Args() []string {
    var args = []string{"db", "prefix"}
    return args
}

