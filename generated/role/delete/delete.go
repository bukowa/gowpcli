package delete

//Deletes an existing role.

type Delete struct {
    
    RoleKey string // <role-key>
    
}

//## OPTIONS
//
//<role-key>
//: The internal name of the role.
//
//## EXAMPLES
//
//    # Delete approver role.
//    $ wp role delete approver
//    Success: Role with key 'approver' deleted.
//
//    # Delete productadmin role.
//    wp role delete productadmin
//    Success: Role with key 'productadmin' deleted.
//
//