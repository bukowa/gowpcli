package incr

//Increments a value in the object cache.

type Incr struct {
    
    Key string // <key>
    
    Offset string // [<offset>]
    
    Group string // [<group>]
    
}

//Errors if the value can't be incremented.
//
//## OPTIONS
//
//<key>
//: Cache key.
//
//[<offset>]
//: The amount by which to increment the item's value.
//---
//default: 1
//---
//
//[<group>]
//: Method for grouping data within the cache which allows the same key to be used across groups.
//---
//default: default
//---
//
//## EXAMPLES
//
//    # Increase cache value.
//    $ wp cache incr my_key 2 my_group
//    50
//
//