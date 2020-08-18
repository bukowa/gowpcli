package list

//List all terms associated with an object.

type List struct {
    
    Id string // <id>
    
    Taxonomy []string // <taxonomy>...
    
    Field string // [--field=<field>]
    
    Fields string // [--fields=<fields>]
    
    Format string // [--format=<format>]
    
}

//<id>
//: ID for the object.
//
//<taxonomy>...
//: One or more taxonomies to list.
//
//[--field=<field>]
//: Prints the value of a single field for each term.
//
//[--fields=<fields>]
//: Limit the output to specific row fields.
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
//  - count
//  - ids
//---
//
//## AVAILABLE FIELDS
//
//These fields will be displayed by default for each term:
//
//* term_id
//* name
//* slug
//* taxonomy
//
//These fields are optionally available:
//
//* term_taxonomy_id
//* description
//* term_group
//* parent
//* count
//
//