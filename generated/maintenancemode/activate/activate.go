package activate

//Activates maintenance mode.

type Activate struct {
    
    Force bool // [--force]
    
}

//[--force]
//: Force maintenance mode activation operation.
//
//## EXAMPLES
//
//    $ wp maintenance-mode activate
//    Enabling Maintenance mode...
//    Success: Activated Maintenance mode.
//
//