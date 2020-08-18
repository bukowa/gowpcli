package spam

//Marks one or more users as spam.

type Spam struct {
    
    Id []string // <id>...
    
}

//## OPTIONS
//
//<id>...
//: One or more IDs of users to mark as spam.
//
//## EXAMPLES
//
//    $ wp user spam 123
//    User 123 marked as spam.
//    Success: Spamed 1 of 1 users.
//
//