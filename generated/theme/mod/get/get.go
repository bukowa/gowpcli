package get

//Gets one or more theme mods.

type Get struct {
    
    Mod []string // [<mod>...]
    
    Field string // [--field=<field>]
    
    All bool // [--all]
    
    Format string // [--format=<format>]
    
}

//## OPTIONS
//
//[<mod>...]
//: One or more mods to get.
//
//[--field=<field>]
//: Returns the value of a single field.
//
//[--all]
//: List all theme mods
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
//    # Get all theme mods.
//    $ wp theme mod get --all
//    +------------------+---------+
//    | key              | value   |
//    +------------------+---------+
//    | background_color | dd3333  |
//    | link_color       | #dd9933 |
//    | main_text_color  | #8224e3 |
//    +------------------+---------+
//
//    # Get single theme mod in JSON format.
//    $ wp theme mod get background_color --format=json
//    [{"key":"background_color","value":"dd3333"}]
//
//    # Get value of a single theme mod.
//    $ wp theme mod get background_color --field=value
//    dd3333
//
//    # Get multiple theme mods.
//    $ wp theme mod get background_color header_textcolor
//    +------------------+--------+
//    | key              | value  |
//    +------------------+--------+
//    | background_color | dd3333 |
//    | header_textcolor |        |
//    +------------------+--------+
//
//