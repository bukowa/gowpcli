package list

//List all metadata associated with an object.

type List struct {
    
    Id string // <id>
    
    Keys string // [--keys=<keys>]
    
    Fields string // [--fields=<fields>]
    
    Format string // [--format=<format>]
    
    Orderby string // [--orderby=<fields>]
    
    Order string // [--order=<order>]
    
    Unserialize bool // [--unserialize]
    
}

//## OPTIONS
//
//<id>
//: ID for the object.
//
//[--keys=<keys>]
//: Limit output to metadata of specific keys.
//
//[--fields=<fields>]
//: Limit the output to specific row fields. Defaults to id,meta_key,meta_value.
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
//  - count
//---
//
//[--orderby=<fields>]
//: Set orderby which field.
//---
//default: id
//options:
// - id
// - meta_key
// - meta_value
//---
//
//[--order=<order>]
//: Set ascending or descending order.
//---
//default: asc
//options:
// - asc
// - desc
//---
//
//[--unserialize]
//: Unserialize meta_value output.
//
//