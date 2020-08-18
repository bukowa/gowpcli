package get

//Gets a transient value.

type Get struct {
    
    Key string // <key>
    
    Format string // [--format=<format>]
    
    Network bool // [--network]
    
}

//For a more complete explanation of the transient cache, including the
//network|site cache, please see docs for `wp transient`.
//
//## OPTIONS
//
//<key>
//: Key for the transient.
//
//[--format=<format>]
//: Render output in a particular format.
//---
//default: table
//options:
//  - table
//  - csv
//  - json
//  - yaml
//---
//
//[--network]
//: Get the value of a network|site transient. On single site, this is
//is a specially-named cache key. On multisite, this is a global cache
//(instead of local to the site).
//
//## EXAMPLES
//
//    $ wp transient get sample_key
//    test data
//
//    $ wp transient get random_key
//    Warning: Transient with key "random_key" is not set.
//
//