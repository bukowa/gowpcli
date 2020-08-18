package list

//Lists available WP-CLI aliases.

type List struct {
    
    Format string // [--format=<format>]
    
}

//## OPTIONS
//
//[--format=<format>]
//: Render output in a particular format.
//---
//default: yaml
//options:
//  - yaml
//  - json
//  - var_export
//---
//
//## EXAMPLES
//
//    # List all available aliases.
//    $ wp cli alias list
//    ---
//    @all: Run command against every registered alias.
//    @prod:
//      ssh: runcommand@runcommand.io~/webapps/production
//    @dev:
//      ssh: vagrant@192.168.50.10/srv/www/runcommand.dev
//    @both:
//      - @prod
//      - @dev
//
//