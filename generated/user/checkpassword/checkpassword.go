package checkpassword

//Checks if a user's password is valid or not.

type CheckPassword struct {
    
    User string // <user>
    
    UserPass string // <user_pass>
    
}

//## OPTIONS
//
//<user>
//: The user login, user email or user ID of the user to check credentials for.
//
//<user_pass>
//: A string that contains the plain text password for the user.
//
//## EXAMPLES
//
//    # Check whether given credentials are valid; exit status 0 if valid, otherwise 1
//    $ wp user check-password admin adminpass
//    $ echo $?
//    1
//
//    # Bash script for checking whether given credentials are valid or not
//    if ! $(wp user check-password admin adminpass); then
//     notify-send "Invalid Credentials";
//    fi
//
//