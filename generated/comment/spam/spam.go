package spam

//Marks a comment as spam.

type Spam struct {
    
    Id []string // <id>...
    
}

//## OPTIONS
//
//<id>...
//: The IDs of the comments to mark as spam.
//
//## EXAMPLES
//
//    # Spam comment.
//    $ wp comment spam 1337
//    Success: Marked as spam comment 1337.
//
//