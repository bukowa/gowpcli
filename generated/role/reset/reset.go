package reset

//Resets any default role to default capabilities.

type Reset struct {
    
    RoleKey []string // [<role-key>...]
    
    All bool // [--all]
    
}

//## OPTIONS
//
//[<role-key>...]
//: The internal name of one or more roles to reset.
//
//[--all]
//: If set, all default roles will be reset.
//
//## EXAMPLES
//
//    # Reset role.
//    $ wp role reset administrator author contributor
//    Success: Reset 1/3 roles.
//
//    # Reset all default roles.
//    $ wp role reset --all
//    Success: All default roles reset.
//
//