package approve

//Approves a comment.

type Approve struct {
    
    Id []string // <id>...
    
}

//## OPTIONS
//
//<id>...
//: The IDs of the comments to approve.
//
//## EXAMPLES
//
//    # Approve comment.
//    $ wp comment approve 1337
//    Success: Approved comment 1337.
//
//