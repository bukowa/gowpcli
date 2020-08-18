/*
## EXAMPLES
	    # Install the Dutch core language pack.
	    $ wp language core install nl_NL
	    Success: Language installed.
	    # Activate the Dutch core language pack.
	    $ wp language core activate nl_NL
	    Success: Language activated.
	    # Install the Dutch language pack for Twenty Seventeen.
	    $ wp language theme install twentyseventeen nl_NL
	    Success: Language installed.
	    # Install the Dutch language pack for Akismet.
	    $ wp language plugin install akismet nl_NL
	    Success: Language installed.
	
 */
package language


// Language //Installs, activates, and manages language packs.
type Language struct {
}

func (l Language) Args() []string {
    var args = []string{"language"}
    return args
}

