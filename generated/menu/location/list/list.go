package list

//Lists locations for the current theme.

type List struct {
    
    Format string // [--format=<format>]
    
}

//## OPTIONS
//
//[--format=<format>]
//: Render output in a particular format.
//---
//default: table
//options:
//  - table
//  - csv
//  - json
//  - count
//  - yaml
//  - ids
//---
//
//## AVAILABLE FIELDS
//
//These fields will be displayed by default for each location:
//
//* name
//* description
//
//## EXAMPLES
//
//    $ wp menu location list
//    +----------+-------------------+
//    | location | description       |
//    +----------+-------------------+
//    | primary  | Primary Menu      |
//    | social   | Social Links Menu |
//    +----------+-------------------+
//
//