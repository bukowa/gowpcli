/*
## EXAMPLES
	    # Create a POT file for the WordPress plugin/theme in the current directory
	    $ wp i18n make-pot . languages/my-plugin.pot
	
 */
package i18n


// I18n //Provides internationalization tools for WordPress projects.
type I18n struct {
}

func (i I18n) Args() []string {
    var args = []string{"i18n"}
    return args
}

