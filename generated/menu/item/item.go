/*
## EXAMPLES
	    # Add an existing post to an existing menu
	    $ wp menu item add-post sidebar-menu 33 --title="Custom Test Post"
	    Success: Menu item added.
	    # Create a new menu link item
	    $ wp menu item add-custom sidebar-menu Apple http://apple.com
	    Success: Menu item added.
	    # Delete menu item
	    $ wp menu item delete 45
	    Success: 1 menu item deleted.
	
	
 */
package item


// Item //List, add, and delete items associated with a menu.
type Item struct {
}

func (i Item) Args() []string {
    var args = []string{"menu", "item"}
    return args
}

