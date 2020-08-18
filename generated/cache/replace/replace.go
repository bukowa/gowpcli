package replace

//Replaces a value in the object cache, if the value already exists.

type Replace struct {
    
    Key string // <key>
    
    Value string // <value>
    
    Group string // [<group>]
    
    Expiration string // [<expiration>]
    
}

//Errors if the value can't be replaced.
//
//## OPTIONS
//
//<key>
//: Cache key.
//
//<value>
//: Value to replace.
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
//    # Replace cache.
//    $ wp cache replace my_key new_value my_group
//    Success: Replaced object 'my_key' in group 'my_group'.
//
//