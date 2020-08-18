package unapprove

//Unapproves a comment.

type Unapprove struct {
    
    Id []string // <id>...
    
}

//## OPTIONS
//
//<id>...
//: The IDs of the comments to unapprove.
//
//## EXAMPLES
//
//    # Unapprove comment.
//    $ wp comment unapprove 1337
//    Success: Unapproved comment 1337.
//
//