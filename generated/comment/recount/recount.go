package recount

//Recalculates the comment_count value for one or more posts.

type Recount struct {
    
    Id []string // <id>...
    
}

//## OPTIONS
//
//<id>...
//: IDs for one or more posts to update.
//
//## EXAMPLES
//
//    # Recount comment for the post.
//    $ wp comment recount 123
//    Updated post 123 comment count to 67.
//
//