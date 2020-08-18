package remove

//Removes capabilities from a given role.

type Remove struct {
    
    Role string // <role>
    
    Cap []string // <cap>...
    
}

//## OPTIONS
//
//<role>
//: Key for the role.
//
//<cap>...
//: One or more capabilities to remove.
//
//## EXAMPLES
//
//    # Remove 'spectate' capability from 'author' role.
//    $ wp cap remove author spectate
//    Success: Removed 1 capability from 'author' role.
//
//