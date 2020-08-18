package removerole

//Removes a user's role.

type RemoveRole struct {
    
    User string // <user>
    
    Role string // [<role>]
    
}

//## OPTIONS
//
//<user>
//: User ID, user email, or user login.
//
//[<role>]
//: A specific role to remove.
//
//## EXAMPLES
//
//    $ wp user remove-role 12 author
//    Success: Removed 'author' role for johndoe (12).
//
//