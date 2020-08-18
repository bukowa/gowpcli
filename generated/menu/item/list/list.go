package list

//Gets a list of items associated with a menu.

type List struct {
    
    Menu string // <menu>
    
    Fields string // [--fields=<fields>]
    
    Format string // [--format=<format>]
    
}

//## OPTIONS
//
//<menu>
//: The name, slug, or term ID for the menu.
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
//  - count
//  - ids
//  - yaml
//---
//
//## AVAILABLE FIELDS
//
//These fields will be displayed by default for each menu item:
//
//* db_id
//* type
//* title
//* link
//* position
//
//These fields are optionally available:
//
//* menu_item_parent
//* object_id
//* object
//* type
//* type_label
//* target
//* attr_title
//* description
//* classes
//* xfn
//
//## EXAMPLES
//
//    $ wp menu item list main-menu
//    +-------+-----------+-------------+---------------------------------+----------+
//    | db_id | type      | title       | link                            | position |
//    +-------+-----------+-------------+---------------------------------+----------+
//    | 5     | custom    | Home        | http://example.com              | 1        |
//    | 6     | post_type | Sample Page | http://example.com/sample-page/ | 2        |
//    +-------+-----------+-------------+---------------------------------+----------+
//
//