package get

//Gets details about a registered post type.

type Get struct {
    
    PostType string // <post-type>
    
    Field string // [--field=<field>]
    
    Fields string // [--fields=<fields>]
    
    Format string // [--format=<format>]
    
}

//## OPTIONS
//
//<post-type>
//: Post type slug
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
//These fields will be displayed by default for the specified post type:
//
//* name
//* label
//* description
//* hierarchical
//* public
//* capability_type
//* labels
//* cap
//* supports
//
//These fields are optionally available:
//
//* count
//
//## EXAMPLES
//
//    # Get details about the 'page' post type.
//    $ wp post-type get page --fields=name,label,hierarchical --format=json
//    {"name":"page","label":"Pages","hierarchical":true}
//
//