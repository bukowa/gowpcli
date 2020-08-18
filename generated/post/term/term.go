/*
## EXAMPLES
	    # Set post terms
	    $ wp post term set 123 test category
	    Success: Set terms.
	
	
 */
package term


// Term //Adds, updates, removes, and lists post terms.
type Term struct {
}

func (t Term) Args() []string {
    var args = []string{"post", "term"}
    return args
}

