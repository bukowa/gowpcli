package setrole

//Sets the user role.

type SetRole struct {
    
    User string // <user>
    
    Role string // [<role>]
    
}

//## OPTIONS
//
//<user>
//: User ID, user email, or user login.
//
//[<role>]
//: Make the user have the specified role. If not passed, the default role is
//used.
//
//## EXAMPLES
//
//    $ wp user set-role 12 author
//    Success: Added johndoe (12) to http://example.com as author.
//
//