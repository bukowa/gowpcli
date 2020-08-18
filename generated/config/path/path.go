/*
## EXAMPLES
	    # Get wp-config.php file path
	    $ wp config path
	    /home/person/htdocs/project/wp-config.php
	
 */
package path


// Path //Gets the path to wp-config.php file.
type Path struct {
}

func (p Path) Args() []string {
    var args = []string{"config", "path"}
    return args
}

