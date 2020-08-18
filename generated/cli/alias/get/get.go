package get

//Gets the value for an alias.

type Get struct {
    
    Key string // <key>
    
}

//## OPTIONS
//
//<key>
//: Key for the alias.
//
//## EXAMPLES
//
//    # Get alias.
//    $ wp cli alias get @prod
//    ssh: dev@somedeve.env:12345/home/dev/
//
//