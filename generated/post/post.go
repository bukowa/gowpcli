/*
## EXAMPLES
	    # Create a new post.
	    $ wp post create --post_type=post --post_title='A sample post'
	    Success: Created post 123.
	    # Update an existing post.
	    $ wp post update 123 --post_status=draft
	    Success: Updated post 123.
	    # Delete an existing post.
	    $ wp post delete 123
	    Success: Trashed post 123.
	
 */
package post


// Post //Manages posts, content, and meta.
type Post struct {
}

func (p Post) Args() []string {
    var args = []string{"post"}
    return args
}

