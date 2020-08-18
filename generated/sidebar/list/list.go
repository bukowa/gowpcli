package list

//Lists registered sidebars.

type List struct {
    
    Fields string // [--fields=<fields>]
    
    Format string // [--format=<format>]
    
}

//## OPTIONS
//
//[--fields=<fields>]
//: Limit the output to specific object fields.
//
//[--format=<format>]
//: Render output in a particular format.
//---
//default: table
//options:
//  - table
//  - csv
//  - json
//  - ids
//  - count
//  - yaml
//---
//
//## AVAILABLE FIELDS
//
//These fields will be displayed by default for each sidebar:
//
//* name
//* id
//* description
//
//These fields are optionally available:
//
//* class
//* before_widget
//* after_widget
//* before_title
//* after_title
//
//## EXAMPLES
//
//    $ wp sidebar list --fields=name,id --format=csv
//    name,id
//    "Widget Area",sidebar-1
//    "Inactive Widgets",wp_inactive_widgets
//
//