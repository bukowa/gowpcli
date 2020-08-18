package delete

//Deletes a meta field.

type Delete struct {
    
    User string // <user>
    
    Key string // <key>
    
    Value string // [<value>]
    
}

//## OPTIONS
//
//<user>
//: The user login, user email, or user ID of the user to delete metadata from.
//
//<key>
//: The metadata key.
//
//[<value>]
//: The value to delete. If omitted, all rows with key will deleted.
//
//## EXAMPLES
//
//    # Delete user meta
//    $ wp user meta delete 123 bio
//    Success: Deleted custom field.
//
//