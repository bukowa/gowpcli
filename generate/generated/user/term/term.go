/*
## EXAMPLES
	    # Set user terms
	    $ wp user term set 123 test category
	    Success: Set terms.
	
	
 */
package term


// Term //Adds, updates, removes, and lists user terms.
type Term struct {
}

func (t Term) Args() []string {
    var args = []string{"user", "term"}
    return args
}

