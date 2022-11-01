/*
## EXAMPLES
	    # Enable the auto-updates for a theme
	    $ wp theme auto-updates enable twentysixteen
	    Theme auto-updates for 'twentysixteen' enabled.
	    Success: Enabled 1 of 1 theme auto-updates.
	    # Disable the auto-updates for a theme
	    $ wp theme auto-updates disable twentysixteen
	    Theme auto-updates for 'twentysixteen' disabled.
	    Success: Disabled 1 of 1 theme auto-updates.
	    # Get the status of theme auto-updates
	    $ wp theme auto-updates status twentysixteen
	    Auto-updates for theme 'twentysixteen' are disabled.
	
	
 */
package autoupdates


// AutoUpdates //Manages theme auto-updates.
type AutoUpdates struct {
}

func (a AutoUpdates) Args() []string {
    var args = []string{"theme", "auto-updates"}
    return args
}

