package list

//Lists transients and their values.

type List struct {
    
    Search string // [--search=<pattern>]
    
    Exclude string // [--exclude=<pattern>]
    
    Network bool // [--network]
    
    Unserialize bool // [--unserialize]
    
    HumanReadable bool // [--human-readable]
    
    Fields string // [--fields=<fields>]
    
    Format string // [--format=<format>]
    
}

//## OPTIONS
//
//[--search=<pattern>]
//: Use wildcards ( * and ? ) to match transient name.
//
//[--exclude=<pattern>]
//: Pattern to exclude. Use wildcards ( * and ? ) to match transient name.
//
//[--network]
//: Get the values of network|site transients. On single site, this is
//a specially-named cache key. On multisite, this is a global cache
//(instead of local to the site).
//
//[--unserialize]
//: Unserialize transient values in output.
//
//[--human-readable]
//: Human-readable output for expirations.
//
//[--fields=<fields>]
//: Limit the output to specific object fields.
//
//[--format=<format>]
//: The serialization format for the value.
//---
//default: table
//options:
//  - table
//  - json
//  - csv
//  - count
//  - yaml
//---
//
//## AVAILABLE FIELDS
//
//This field will be displayed by default for each matching option:
//
//* name
//* value
//* expiration
//
//## EXAMPLES
//
//    # List all transients
//    $ wp transient list
//     +------+-------+---------------+
//     | name | value | expiration    |
//     +------+-------+---------------+
//     | foo  | bar   | 39 mins       |
//     | foo2 | bar2  | no expiration |
//     | foo3 | bar2  | expired       |
//     | foo4 | bar4  | 4 hours       |
//     +------+-------+---------------+
//
//