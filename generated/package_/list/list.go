package list

//Lists installed WP-CLI packages.

type List struct {
    
    Fields string // [--fields=<fields>]
    
    Format string // [--format=<format>]
    
}

//## OPTIONS
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
//  - ids
//  - json
//  - yaml
//---
//
//## AVAILABLE FIELDS
//
//These fields will be displayed by default for each package:
//
//* name
//* authors
//* version
//* update
//* update_version
//
//These fields are optionally available:
//
//* description
//
//## EXAMPLES
//
//    $ wp package list
//    +-----------------------+------------------------------------------+---------+------------+
//    | name                  | description                              | authors | version    |
//    +-----------------------+------------------------------------------+---------+------------+
//    | wp-cli/server-command | Start a development server for WordPress |         | dev-master |
//    +-----------------------+------------------------------------------+---------+------------+
//
//