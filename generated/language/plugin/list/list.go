package list

//Lists all available languages for one or more plugins.

type List struct {
    
    Plugin []string // [<plugin>...]
    
    All bool // [--all]
    
    Field string // [--field=<field>]
    
    FieldMap map[string]string // [--<field>=<value>]
    
    Fields string // [--fields=<fields>]
    
    Format string // [--format=<format>]
    
}

//## OPTIONS
//
//[<plugin>...]
//: One or more plugins to list languages for.
//
//[--all]
//: If set, available languages for all plugins will be listed.
//
//[--field=<field>]
//: Display the value of a single field.
//
//[--<field>=<value>]
//: Filter results by key=value pairs.
//
//[--fields=<fields>]
//: Limit the output to specific fields.
//
//[--format=<format>]
//: Render output in a particular format.
//---
//default: table
//options:
//  - table
//  - csv
//  - json
//---
//
//## AVAILABLE FIELDS
//
//These fields will be displayed by default for each translation:
//
//* plugin
//* language
//* english_name
//* native_name
//* status
//* update
//* updated
//
//## EXAMPLES
//
//    # List language,english_name,status fields of available languages.
//    $ wp language plugin list --fields=language,english_name,status
//    +----------------+-------------------------+-------------+
//    | language       | english_name            | status      |
//    +----------------+-------------------------+-------------+
//    | ar             | Arabic                  | uninstalled |
//    | ary            | Moroccan Arabic         | uninstalled |
//    | az             | Azerbaijani             | uninstalled |
//
//