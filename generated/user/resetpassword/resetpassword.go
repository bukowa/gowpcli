package resetpassword

//Resets the password for one or more users.

type ResetPassword struct {
    
    User []string // <user>...
    
    SkipEmail bool // [--skip-email]
    
}

//## OPTIONS
//
//<user>...
//: one or more user logins or IDs.
//
//[--skip-email]
//: Don't send an email notification to the affected user(s).
//
//## EXAMPLES
//
//    # Reset the password for two users and send them the change email.
//    $ wp user reset-password admin editor
//    Reset password for admin.
//    Reset password for editor.
//    Success: Passwords reset for 2 users.
//
//