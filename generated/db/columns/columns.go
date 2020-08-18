package columns

//Displays information about a given table.

type Columns struct {
    
    Table string // [<table>]
    
    Format bool // [--format]
    
}

//## OPTIONS
//
//[<table>]
//: Name of the database table.
//
//[--format]
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
//    $ wp db columns wp_posts
//    +-----------------------+---------------------+------+-----+---------------------+----------------+
//    |         Field         |        Type         | Null | Key |       Default       |     Extra      |
//    +-----------------------+---------------------+------+-----+---------------------+----------------+
//    | ID                    | bigint(20) unsigned | NO   | PRI |                     | auto_increment |
//    | post_author           | bigint(20) unsigned | NO   | MUL | 0                   |                |
//    | post_date             | datetime            | NO   |     | 0000-00-00 00:00:00 |                |
//    | post_date_gmt         | datetime            | NO   |     | 0000-00-00 00:00:00 |                |
//    | post_content          | longtext            | NO   |     |                     |                |
//    | post_title            | text                | NO   |     |                     |                |
//    | post_excerpt          | text                | NO   |     |                     |                |
//    | post_status           | varchar(20)         | NO   |     | publish             |                |
//    | comment_status        | varchar(20)         | NO   |     | open                |                |
//    | ping_status           | varchar(20)         | NO   |     | open                |                |
//    | post_password         | varchar(255)        | NO   |     |                     |                |
//    | post_name             | varchar(200)        | NO   | MUL |                     |                |
//    | to_ping               | text                | NO   |     |                     |                |
//    | pinged                | text                | NO   |     |                     |                |
//    | post_modified         | datetime            | NO   |     | 0000-00-00 00:00:00 |                |
//    | post_modified_gmt     | datetime            | NO   |     | 0000-00-00 00:00:00 |                |
//    | post_content_filtered | longtext            | NO   |     |                     |                |
//    | post_parent           | bigint(20) unsigned | NO   | MUL | 0                   |                |
//    | guid                  | varchar(255)        | NO   |     |                     |                |
//    | menu_order            | int(11)             | NO   |     | 0                   |                |
//    | post_type             | varchar(20)         | NO   | MUL | post                |                |
//    | post_mime_type        | varchar(100)        | NO   |     |                     |                |
//    | comment_count         | bigint(20)          | NO   |     | 0                   |                |
//    +-----------------------+---------------------+------+-----+---------------------+----------------+
//
//