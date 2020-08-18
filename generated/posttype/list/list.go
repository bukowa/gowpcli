package list

//Lists registered post types.

type List struct {
    
    FieldMap map[string]string // [--<field>=<value>]
    
    Field string // [--field=<field>]
    
    Fields string // [--fields=<fields>]
    
    Format string // [--format=<format>]
    
}

//## OPTIONS
//
//[--<field>=<value>]
//: Filter by one or more fields (see get_post_types() first parameter for a list of available fields).
//
//[--field=<field>]
//: Prints the value of a single field for each post type.
//
//[--fields=<fields>]
//: Limit the output to specific post type fields.
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
//---
//
//## AVAILABLE FIELDS
//
//These fields will be displayed by default for each post type:
//
//* name
//* label
//* description
//* hierarchical
//* public
//* capability_type
//
//These fields are optionally available:
//
//* count
//
//## EXAMPLES
//
//    # List registered post types
//    $ wp post-type list --format=csv
//    name,label,description,hierarchical,public,capability_type
//    post,Posts,,,1,post
//    page,Pages,,1,1,page
//    attachment,Media,,,1,post
//    revision,Revisions,,,,post
//    nav_menu_item,"Navigation Menu Items",,,,post
//
//    # List post types with 'post' capability type
//    $ wp post-type list --capability_type=post --fields=name,public
//    +---------------+--------+
//    | name          | public |
//    +---------------+--------+
//    | post          | 1      |
//    | attachment    | 1      |
//    | revision      |        |
//    | nav_menu_item |        |
//    +---------------+--------+
//
//