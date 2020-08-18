package get

//Gets details about a registered taxonomy.

type Get struct {
    
    Taxonomy string // <taxonomy>
    
    Field string // [--field=<field>]
    
    Fields string // [--fields=<fields>]
    
    Format string // [--format=<format>]
    
}

//## OPTIONS
//
//<taxonomy>
//: Taxonomy slug.
//
//[--field=<field>]
//: Instead of returning the whole taxonomy, returns the value of a single field.
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
//## AVAILABLE FIELDS
//
//These fields will be displayed by default for the specified taxonomy:
//
//* name
//* label
//* description
//* object_type
//* show_tagcloud
//* hierarchical
//* public
//* labels
//* cap
//
//These fields are optionally available:
//
//* count
//
//## EXAMPLES
//
//    # Get details of `category` taxonomy.
//    $ wp taxonomy get category --fields=name,label,object_type
//    +-------------+------------+
//    | Field       | Value      |
//    +-------------+------------+
//    | name        | category   |
//    | label       | Categories |
//    | object_type | ["post"]   |
//    +-------------+------------+
//
//    # Get capabilities of 'post_tag' taxonomy.
//    $ wp taxonomy get post_tag --field=cap
//    {"manage_terms":"manage_categories","edit_terms":"manage_categories","delete_terms":"manage_categories","assign_terms":"edit_posts"}
//
//