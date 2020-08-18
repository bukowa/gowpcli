/*
## INFO
	Display posts based on all arguments supported by
	[WP_Query()](https://developer.wordpress.org/reference/classes/wp_query/).
## OPTIONS
	[--<field>=<value>]
	: One or more args to pass to WP_Query.
	[--field=<field>]
	: Prints the value of a single field for each post.
	[--fields=<fields>]
	: Limit the output to specific object fields.
	[--format=<format>]
	: Render output in a particular format.
	---
	default: table
	options:
	  - table
	  - csv
	  - ids
	  - json
	  - count
	  - yaml
	---
## AVAILABLE FIELDS
	These fields will be displayed by default for each post:
	* ID
	* post_title
	* post_name
	* post_date
	* post_status
	These fields are optionally available:
	* post_author
	* post_date_gmt
	* post_content
	* post_excerpt
	* comment_status
	* ping_status
	* post_password
	* to_ping
	* pinged
	* post_modified
	* post_modified_gmt
	* post_content_filtered
	* post_parent
	* guid
	* menu_order
	* post_type
	* post_mime_type
	* comment_count
	* filter
	* url
## EXAMPLES
	    # List post
	    $ wp post list --field=ID
	    568
	    829
	    1329
	    1695
	    # List posts in JSON
	    $ wp post list --post_type=post --posts_per_page=5 --format=json
	    [{"ID":1,"post_title":"Hello world!","post_name":"hello-world","post_date":"2015-06-20 09:00:10","post_status":"publish"},{"ID":1178,"post_title":"Markup: HTML Tags and Formatting","post_name":"markup-html-tags-and-formatting","post_date":"2013-01-11 20:22:19","post_status":"draft"}]
	    # List all pages
	    $ wp post list --post_type=page --fields=post_title,post_status
	    +-------------+-------------+
	    | post_title  | post_status |
	    +-------------+-------------+
	    | Sample Page | publish     |
	    +-------------+-------------+
	    # List ids of all pages and posts
	    $ wp post list --post_type=page,post --format=ids
	    15 25 34 37 198
	    # List given posts
	    $ wp post list --post__in=1,3
	    +----+--------------+-------------+---------------------+-------------+
	    | ID | post_title   | post_name   | post_date           | post_status |
	    +----+--------------+-------------+---------------------+-------------+
	    | 3  | Lorem Ipsum  | lorem-ipsum | 2016-06-01 14:34:36 | publish     |
	    | 1  | Hello world! | hello-world | 2016-06-01 14:31:12 | publish     |
	    +----+--------------+-------------+---------------------+-------------+
	
 */
package list
import utils "github.com/bukowa/gowpcli"

// List //Gets a list of posts.
type List struct {
    FieldMap map[string]string // [--<field>=<value>]
    Field string // [--field=<field>]
    Fields string // [--fields=<fields>]
    Format string // [--format=<format>]
}

func (l List) Args() []string {
    var args = []string{"post", "list"}
    args = utils.MakeArg(args, "[--<field>=<value>]", l.FieldMap)
    args = utils.MakeArg(args, "[--field=<field>]", l.Field)
    args = utils.MakeArg(args, "[--fields=<fields>]", l.Fields)
    args = utils.MakeArg(args, "[--format=<format>]", l.Format)
    return args
}

