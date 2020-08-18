package untrash

//Untrashes a comment.

type Untrash struct {
    
    Id []string // <id>...
    
}

//## OPTIONS
//
//<id>...
//: The IDs of the comments to untrash.
//
//## EXAMPLES
//
//    # Untrash comment.
//    $ wp comment untrash 1337
//    Success: Untrashed comment 1337.
//
//