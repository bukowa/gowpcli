package addrole

//Adds a role for a user.

type AddRole struct {
    
    User string // <user>
    
    Role string // <role>
    
}

//## OPTIONS
//
//<user>
//: User ID, user email, or user login.
//
//<role>
//: Add the specified role to the user.
//
//## EXAMPLES
//
//    $ wp user add-role 12 author
//    Success: Added 'author' role for johndoe (12).
//
//