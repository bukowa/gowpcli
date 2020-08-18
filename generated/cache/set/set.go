package set

//Sets a value to the object cache, regardless of whether it already exists.

type Set struct {
    
    Key string // <key>
    
    Value string // <value>
    
    Group string // [<group>]
    
    Expiration string // [<expiration>]
    
}

//Errors if the value can't be set.
//
//## OPTIONS
//
//<key>
//: Cache key.
//
//<value>
//: Value to set on the key.
//
//[<group>]
//: Method for grouping data within the cache which allows the same key to be used across groups.
//---
//default: default
//---
//
//[<expiration>]
//: Define how long to keep the value, in seconds. `0` means as long as possible.
//---
//default: 0
//---
//
//## EXAMPLES
//
//    # Set cache.
//    $ wp cache set my_key my_value my_group 300
//    Success: Set object 'my_key' in group 'my_group'.
//
//