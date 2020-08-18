package update

//Updates one or more comments.

type Update struct {
    
    Id []string // <id>...
    
    Field string // --<field>=<value>
    
}

//## OPTIONS
//
//<id>...
//: One or more IDs of comments to update.
//
//--<field>=<value>
//: One or more fields to update. See wp_update_comment().
//
//## EXAMPLES
//
//    # Update comment.
//    $ wp comment update 123 --comment_author='That Guy'
//    Success: Updated comment 123.
//
//