package exists

//Verifies whether a post exists.

type Exists struct {
    
    Id string // <id>
    
}

//Displays a success message if the post does exist.
//
//## OPTIONS
//
//<id>
//: The ID of the post to check.
//
//## EXAMPLES
//
//    # The post exists.
//    $ wp post exists 1
//    Success: Post with ID 1337 exists.
//    $ echo $?
//    0
//
//    # The post does not exist.
//    $ wp post exists 10000
//    $ echo $?
//    1
//
//