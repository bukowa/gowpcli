package add

//Grants super admin privileges to one or more users.

type Add struct {
    
    User []string // <user>...
    
}

//## OPTIONS
//
//<user>...
//: One or more user IDs, user emails, or user logins.
//
//## EXAMPLES
//
//    $ wp super-admin add superadmin2
//    Success: Granted super-admin capabilities.
//
//