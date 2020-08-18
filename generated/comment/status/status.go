package status

//Gets the status of a comment.

type Status struct {
    
    Id string // <id>
    
}

//## OPTIONS
//
//<id>
//: The ID of the comment to check.
//
//## EXAMPLES
//
//    # Get status of comment.
//    $ wp comment status 1337
//    approved
//
//