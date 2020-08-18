package exists

//Checks if a role exists.

type Exists struct {
    
    RoleKey string // <role-key>
    
}

//Exits with return code 0 if the role exists, 1 if it does not.
//
//## OPTIONS
//
//<role-key>
//: The internal name of the role.
//
//## EXAMPLES
//
//    # Check if a role exists.
//    $ wp role exists editor
//    Success: Role with ID 'editor' exists.
//
//