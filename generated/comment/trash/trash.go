package trash

//Trashes a comment.

type Trash struct {
    
    Id []string // <id>...
    
}

//## OPTIONS
//
//<id>...
//: The IDs of the comments to trash.
//
//## EXAMPLES
//
//    # Trash comment.
//    $ wp comment trash 1337
//    Success: Trashed comment 1337.
//
//