package delete

//Deletes a comment.

type Delete struct {
    
    Id []string // <id>...
    
    Force bool // [--force]
    
}

//## OPTIONS
//
//<id>...
//: One or more IDs of comments to delete.
//
//[--force]
//: Skip the trash bin.
//
//## EXAMPLES
//
//    # Delete comment.
//    $ wp comment delete 1337 --force
//    Success: Deleted comment 1337.
//
//    # Delete multiple comments.
//    $ wp comment delete 1337 2341 --force
//    Success: Deleted comment 1337.
//    Success: Deleted comment 2341.
//
//