package exists

//Verifies whether a comment exists.

type Exists struct {
    
    Id string // <id>
    
}

//Displays a success message if the comment does exist.
//
//## OPTIONS
//
//<id>
//: The ID of the comment to check.
//
//## EXAMPLES
//
//    # Check whether comment exists.
//    $ wp comment exists 1337
//    Success: Comment with ID 1337 exists.
//
//