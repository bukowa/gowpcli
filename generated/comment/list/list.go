package list

//Gets a list of comments.

type List struct {
    
    FieldMap map[string]string // [--<field>=<value>]
    
    Field string // [--field=<field>]
    
    Fields string // [--fields=<fields>]
    
    Format string // [--format=<format>]
    
}

//Display comments based on all arguments supported by
//[WP_Comment_Query()](https://developer.wordpress.org/reference/classes/WP_Comment_Query/__construct/).
//
//## OPTIONS
//
//[--<field>=<value>]
//: One or more args to pass to WP_Comment_Query.
//
//[--field=<field>]
//: Prints the value of a single field for each comment.
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
//  - ids
//  - csv
//  - json
//  - count
//  - yaml
//---
//
//## AVAILABLE FIELDS
//
//These fields will be displayed by default for each comment:
//
//* comment_ID
//* comment_post_ID
//* comment_date
//* comment_approved
//* comment_author
//* comment_author_email
//
//These fields are optionally available:
//
//* comment_author_url
//* comment_author_IP
//* comment_date_gmt
//* comment_content
//* comment_karma
//* comment_agent
//* comment_type
//* comment_parent
//* user_id
//* url
//
//## EXAMPLES
//
//    # List comment IDs.
//    $ wp comment list --field=ID
//    22
//    23
//    24
//
//    # List comments of a post.
//    $ wp comment list --post_id=1 --fields=ID,comment_date,comment_author
//    +------------+---------------------+----------------+
//    | comment_ID | comment_date        | comment_author |
//    +------------+---------------------+----------------+
//    | 1          | 2015-06-20 09:00:10 | Mr WordPress   |
//    +------------+---------------------+----------------+
//
//    # List approved comments.
//    $ wp comment list --number=3 --status=approve --fields=ID,comment_date,comment_author
//    +------------+---------------------+----------------+
//    | comment_ID | comment_date        | comment_author |
//    +------------+---------------------+----------------+
//    | 1          | 2015-06-20 09:00:10 | Mr WordPress   |
//    | 30         | 2013-03-14 12:35:07 | John Doe       |
//    | 29         | 2013-03-14 11:56:08 | Jane Doe       |
//    +------------+---------------------+----------------+
//
//