/*
## OPTIONS
	[--match=<url>]
	: Show rewrite rules matching a particular URL.
	[--source=<source>]
	: Show rewrite rules from a particular source.
	[--fields=<fields>]
	: Limit the output to specific fields. Defaults to match,query,source.
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
	    $ wp rewrite list --format=csv
	    match,query,source
	    ^wp-json/?$,index.php?rest_route=/,other
	    ^wp-json/(.*)?,index.php?rest_route=/$matches[1],other
	    category/(.+?)/feed/(feed|rdf|rss|rss2|atom)/?$,index.php?category_name=$matches[1]&feed=$matches[2],category
	    category/(.+?)/(feed|rdf|rss|rss2|atom)/?$,index.php?category_name=$matches[1]&feed=$matches[2],category
	    category/(.+?)/embed/?$,index.php?category_name=$matches[1]&embed=true,category
	
 */
package list
import utils "github.com/bukowa/gowpcli"

// List //Gets a list of the current rewrite rules.
type List struct {
    Match string // [--match=<url>]
    Source string // [--source=<source>]
    Fields string // [--fields=<fields>]
    Format string // [--format=<format>]
}

func (l List) Args() []string {
    var args = []string{"rewrite", "list"}
    args = utils.MakeArg(args, "[--match=<url>]", l.Match)
    args = utils.MakeArg(args, "[--source=<source>]", l.Source)
    args = utils.MakeArg(args, "[--fields=<fields>]", l.Fields)
    args = utils.MakeArg(args, "[--format=<format>]", l.Format)
    return args
}

