package imagesize

//Lists image sizes registered with WordPress.

type ImageSize struct {
    
    Fields string // [--fields=<fields>]
    
    Format string // [--format=<format>]
    
}

//## OPTIONS
//
//[--fields=<fields>]
//: Limit the output to specific fields. Defaults to all fields.
//
//[--format=<format>]
//: Render output in a specific format
//---
//default: table
//options:
//  - table
//  - json
//  - csv
//  - yaml
//  - count
//---
//
//## AVAILABLE FIELDS
//
//These fields will be displayed by default for each image size:
//* name
//* width
//* height
//* crop
//* ratio
//
//## EXAMPLES
//
//    # List all registered image sizes
//    $ wp media image-size
//    +---------------------------+-------+--------+-------+-------+
//    | name                      | width | height | crop  | ratio |
//    +---------------------------+-------+--------+-------+-------+
//    | full                      |       |        | N/A   | N/A   |
//    | twentyfourteen-full-width | 1038  | 576    | hard  | 173:96|
//    | large                     | 1024  | 1024   | soft  | N/A   |
//    | medium_large              | 768   | 0      | soft  | N/A   |
//    | medium                    | 300   | 300    | soft  | N/A   |
//    | thumbnail                 | 150   | 150    | hard  | 1:1   |
//    +---------------------------+-------+--------+-------+-------+
//
//