package get

//Gets details about a theme.

type Get struct {
    
    Theme string // <theme>
    
    Field string // [--field=<field>]
    
    Fields string // [--fields=<fields>]
    
    Format string // [--format=<format>]
    
}

//## OPTIONS
//
//<theme>
//: The theme to get.
//
//[--field=<field>]
//: Instead of returning the whole theme, returns the value of a single field.
//
//[--fields=<fields>]
//: Limit the output to specific fields. Defaults to all fields.
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
//## EXAMPLES
//
//    $ wp theme get twentysixteen --fields=name,title,version
//    +---------+----------------+
//    | Field   | Value          |
//    +---------+----------------+
//    | name    | Twenty Sixteen |
//    | title   | Twenty Sixteen |
//    | version | 1.2            |
//    +---------+----------------+
//
//