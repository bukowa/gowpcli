package list

//Gets a list of themes.

type List struct {
    
    FieldMap map[string]string // [--<field>=<value>]
    
    Field string // [--field=<field>]
    
    Fields string // [--fields=<fields>]
    
    Format string // [--format=<format>]
    
}

//## OPTIONS
//
//[--<field>=<value>]
//: Filter results based on the value of a field.
//
//[--field=<field>]
//: Prints the value of a single field for each theme.
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
//  - yaml
//---
//
//## AVAILABLE FIELDS
//
//These fields will be displayed by default for each theme:
//
//* name
//* status
//* update
//* version
//
//These fields are optionally available:
//
//* update_version
//* update_package
//* update_id
//* title
//* description
//
//## EXAMPLES
//
//    # List themes
//    $ wp theme list --status=inactive --format=csv
//    name,status,update,version
//    twentyfourteen,inactive,none,1.7
//    twentysixteen,inactive,available,1.1
//
//