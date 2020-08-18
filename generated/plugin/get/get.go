package get

//Gets details about an installed plugin.

type Get struct {
    
    Plugin string // <plugin>
    
    Field string // [--field=<field>]
    
    Fields string // [--fields=<fields>]
    
    Format string // [--format=<format>]
    
}

//## OPTIONS
//
//<plugin>
//: The plugin to get.
//
//[--field=<field>]
//: Instead of returning the whole plugin, returns the value of a single field.
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
//## EXAMPLES
//
//    $ wp plugin get bbpress --format=json
//    {"name":"bbpress","title":"bbPress","author":"The bbPress Contributors","version":"2.6-alpha","description":"bbPress is forum software with a twist from the creators of WordPress.","status":"active"}
//
//