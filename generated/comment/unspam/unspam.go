package unspam

//Unmarks a comment as spam.

type Unspam struct {
    
    Id []string // <id>...
    
}

//## OPTIONS
//
//<id>...
//: The IDs of the comments to unmark as spam.
//
//## EXAMPLES
//
//    # Unspam comment.
//    $ wp comment unspam 1337
//    Success: Unspammed comment 1337.
//
//