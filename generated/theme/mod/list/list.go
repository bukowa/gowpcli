package list

//Gets a list of theme mods.

type List struct {
    
    Field string // [--field=<field>]
    
    Format string // [--format=<format>]
    
}

//## OPTIONS
//
//[--field=<field>]
//: Returns the value of a single field.
//
//[--format=<format>]
//: Render output in a particular format.
//---
//default: table
//options:
//  - table
//  - json
//  - csv
//  - yaml
//---
//
//## EXAMPLES
//
//    # Gets a list of theme mods.
//    $ wp theme mod list
//    +------------------+---------+
//    | key              | value   |
//    +------------------+---------+
//    | background_color | dd3333  |
//    | link_color       | #dd9933 |
//    | main_text_color  | #8224e3 |
//    +------------------+---------+
//
//