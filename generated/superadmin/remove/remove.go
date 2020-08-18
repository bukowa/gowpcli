package remove

//Removes super admin privileges from one or more users.

type Remove struct {
    
    User []string // <user>...
    
}

//## OPTIONS
//
//<user>...
//: One or more user IDs, user emails, or user logins.
//
//## EXAMPLES
//
//    $ wp super-admin remove superadmin2
//    Success: Revoked super-admin capabilities.
//
//