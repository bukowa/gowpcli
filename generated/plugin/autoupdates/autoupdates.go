/*
## EXAMPLES
	    # Enable the auto-updates for a plugin
	    $ wp plugin auto-updates enable hello
	    Plugin auto-updates for 'hello' enabled.
	    Success: Enabled 1 of 1 plugin auto-updates.
	    # Disable the auto-updates for a plugin
	    $ wp plugin auto-updates disable hello
	    Plugin auto-updates for 'hello' disabled.
	    Success: Disabled 1 of 1 plugin auto-updates.
	    # Get the status of plugin auto-updates
	    $ wp plugin auto-updates status hello
	    Auto-updates for plugin 'hello' are disabled.
	
	
 */
package autoupdates


// AutoUpdates //Manages plugin auto-updates.
type AutoUpdates struct {
}

func (a AutoUpdates) Args() []string {
    var args = []string{"plugin", "auto-updates"}
    return args
}

