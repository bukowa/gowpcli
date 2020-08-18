package unspam

//Removes one or more users from spam.

type Unspam struct {
    
    Id []string // <id>...
    
}

//## OPTIONS
//
//<id>...
//: One or more IDs of users to remove from spam.
//
//## EXAMPLES
//
//    $ wp user unspam 123
//    User 123 removed from spam.
//    Success: Unspamed 1 of 1 users.
//
//