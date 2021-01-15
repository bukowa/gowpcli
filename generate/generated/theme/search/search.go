/*
## INFO
	Displays themes in the WordPress.org theme directory matching a given
	search query.
## OPTIONS
	<search>
	: The string to search for.
	[--page=<page>]
	: Optional page to display.
	---
	default: 1
	---
	[--per-page=<per-page>]
	: Optional number of results to display. Defaults to 10.
	[--field=<field>]
	: Prints the value of a single field for each theme.
	[--fields=<fields>]
	: Ask for specific fields from the API. Defaults to name,slug,author,rating. Acceptable values:
	    **name**: Theme Name
	    **slug**: Theme Slug
	    **version**: Current Version Number
	    **author**: Theme Author
	    **preview_url**: Theme Preview URL
	    **screenshot_url**: Theme Screenshot URL
	    **rating**: Theme Rating
	    **num_ratings**: Number of Theme Ratings
	    **homepage**: Theme Author's Homepage
	    **description**: Theme Description
	    **url**: Theme's URL on wordpress.org
	[--format=<format>]
	: Render output in a particular format.
	---
	default: table
	options:
	  - table
	  - csv
	  - json
	  - count
	  - yaml
	---
## EXAMPLES
	    $ wp theme search photo --per-page=6
	    Success: Showing 6 of 203 themes.
	    +----------------------+----------------------+--------+
	    | name                 | slug                 | rating |
	    +----------------------+----------------------+--------+
	    | Photos               | photos               | 100    |
	    | Infinite Photography | infinite-photography | 100    |
	    | PhotoBook            | photobook            | 100    |
	    | BG Photo Frame       | bg-photo-frame       | 0      |
	    | fPhotography         | fphotography         | 0      |
	    | Photo Perfect        | photo-perfect        | 98     |
	    +----------------------+----------------------+--------+
	
 */
package search
import utils "github.com/bukowa/gowpcli"

// Search //Searches the WordPress.org theme directory.
type Search struct {
    Search string // <search>
    Page string // [--page=<page>]
    PerPage string // [--per-page=<per-page>]
    Field string // [--field=<field>]
    Fields string // [--fields=<fields>]
    Format string // [--format=<format>]
}

func (s Search) Args() []string {
    var args = []string{"theme", "search"}
    args = utils.MakeArg(args, "<search>", s.Search)
    args = utils.MakeArg(args, "[--page=<page>]", s.Page)
    args = utils.MakeArg(args, "[--per-page=<per-page>]", s.PerPage)
    args = utils.MakeArg(args, "[--field=<field>]", s.Field)
    args = utils.MakeArg(args, "[--fields=<fields>]", s.Fields)
    args = utils.MakeArg(args, "[--format=<format>]", s.Format)
    return args
}

