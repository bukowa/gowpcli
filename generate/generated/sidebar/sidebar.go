/*
## INFO
	A [sidebar](https://developer.wordpress.org/themes/functionality/sidebars/) is any widgetized area of your theme.
## EXAMPLES
	    # List sidebars
	    $ wp sidebar list --fields=name,id --format=csv
	    name,id
	    "Widget Area",sidebar-1
	    "Inactive Widgets",wp_inactive_widgets
	
 */
package sidebar


// Sidebar //Lists registered sidebars.
type Sidebar struct {
}

func (s Sidebar) Args() []string {
    var args = []string{"sidebar"}
    return args
}

