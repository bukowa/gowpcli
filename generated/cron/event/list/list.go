package list

//Lists scheduled cron events.

type List struct {
    
    Fields string // [--fields=<fields>]
    
    FieldMap map[string]string // [--<field>=<value>]
    
    Field string // [--field=<field>]
    
    Format string // [--format=<format>]
    
}

//## OPTIONS
//
//[--fields=<fields>]
//: Limit the output to specific object fields.
//
//[--<field>=<value>]
//: Filter by one or more fields.
//
//[--field=<field>]
//: Prints the value of a single field for each event.
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
//These fields will be displayed by default for each cron event:
//* hook
//* next_run_gmt
//* next_run_relative
//* recurrence
//
//These fields are optionally available:
//* time
//* sig
//* args
//* schedule
//* interval
//* next_run
//
//## EXAMPLES
//
//    # List scheduled cron events
//    $ wp cron event list
//    +-------------------+---------------------+---------------------+------------+
//    | hook              | next_run_gmt        | next_run_relative   | recurrence |
//    +-------------------+---------------------+---------------------+------------+
//    | wp_version_check  | 2016-05-31 22:15:13 | 11 hours 57 minutes | 12 hours   |
//    | wp_update_plugins | 2016-05-31 22:15:13 | 11 hours 57 minutes | 12 hours   |
//    | wp_update_themes  | 2016-05-31 22:15:14 | 11 hours 57 minutes | 12 hours   |
//    +-------------------+---------------------+---------------------+------------+
//
//    # List scheduled cron events in JSON
//    $ wp cron event list --fields=hook,next_run --format=json
//    [{"hook":"wp_version_check","next_run":"2016-05-31 10:15:13"},{"hook":"wp_update_plugins","next_run":"2016-05-31 10:15:13"},{"hook":"wp_update_themes","next_run":"2016-05-31 10:15:14"}]
//
//