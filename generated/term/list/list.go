package list

//Lists terms in a taxonomy.

type List struct {
    
    Taxonomy []string // <taxonomy>...
    
    FieldMap map[string]string // [--<field>=<value>]
    
    Field string // [--field=<field>]
    
    Fields string // [--fields=<fields>]
    
    Format string // [--format=<format>]
    
}

//## OPTIONS
//
//<taxonomy>...
//: List terms of one or more taxonomies
//
//[--<field>=<value>]
//: Filter by one or more fields (see get_terms() $args parameter for a list of fields).
//
//[--field=<field>]
//: Prints the value of a single field for each term.
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
//  - ids
//  - json
//  - count
//  - yaml
//---
//
//## AVAILABLE FIELDS
//
//These fields will be displayed by default for each term:
//
//* term_id
//* term_taxonomy_id
//* name
//* slug
//* description
//* parent
//* count
//
//These fields are optionally available:
//
//* url
//
//## EXAMPLES
//
//    # List post categories
//    $ wp term list category --format=csv
//    term_id,term_taxonomy_id,name,slug,description,parent,count
//    2,2,aciform,aciform,,0,1
//    3,3,antiquarianism,antiquarianism,,0,1
//    4,4,arrangement,arrangement,,0,1
//    5,5,asmodeus,asmodeus,,0,1
//
//    # List post tags
//    $ wp term list post_tag --fields=name,slug
//    +-----------+-------------+
//    | name      | slug        |
//    +-----------+-------------+
//    | 8BIT      | 8bit        |
//    | alignment | alignment-2 |
//    | Articles  | articles    |
//    | aside     | aside       |
//    +-----------+-------------+
//
//