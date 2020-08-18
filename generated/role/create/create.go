package create

//Creates a new role.

type Create struct {
    
    RoleKey string // <role-key>
    
    RoleName string // <role-name>
    
    Clone string // [--clone=<role>]
    
}

//## OPTIONS
//
//<role-key>
//: The internal name of the role.
//
//<role-name>
//: The publicly visible name of the role.
//
//[--clone=<role>]
//: Clone capabilities from an existing role.
//
//## EXAMPLES
//
//    # Create role for Approver.
//    $ wp role create approver Approver
//    Success: Role with key 'approver' created.
//
//    # Create role for Product Administrator.
//    $ wp role create productadmin "Product Administrator"
//    Success: Role with key 'productadmin' created.
//
//