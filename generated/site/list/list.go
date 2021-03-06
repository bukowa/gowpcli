/*
## OPTIONS
	[--network=<id>]
	: The network to which the sites belong.
	[--<field>=<value>]
	: Filter by one or more fields (see "Available Fields" section). However,
	'url' isn't an available filter, because it's created from domain + path.
	[--site__in=<value>]
	: Only list the sites with these blog_id values (comma-separated).
	[--field=<field>]
	: Prints the value of a single field for each site.
	[--fields=<fields>]
	: Comma-separated list of fields to show.
	[--format=<format>]
	: Render output in a particular format.
	---
	default: table
	options:
	  - table
	  - csv
	  - count
	  - ids
	  - json
	  - yaml
	---
## AVAILABLE FIELDS
	These fields will be displayed by default for each site:
	* blog_id
	* url
	* last_updated
	* registered
	These fields are optionally available:
	* site_id
	* domain
	* path
	* public
	* archived
	* mature
	* spam
	* deleted
	* lang_id
## EXAMPLES
	    # Output a simple list of site URLs
	    $ wp site list --field=url
	    http://www.example.com/
	    http://www.example.com/subdir/
	
 */
package list
import utils "github.com/bukowa/gowpcli"

// List //Lists all sites in a multisite installation.
type List struct {
    Network string // [--network=<id>]
    FieldMap map[string]string // [--<field>=<value>]
    SiteIn string // [--site__in=<value>]
    Field string // [--field=<field>]
    Fields string // [--fields=<fields>]
    Format string // [--format=<format>]
}

func (l List) Args() []string {
    var args = []string{"site", "list"}
    args = utils.MakeArg(args, "[--network=<id>]", l.Network)
    args = utils.MakeArg(args, "[--<field>=<value>]", l.FieldMap)
    args = utils.MakeArg(args, "[--site__in=<value>]", l.SiteIn)
    args = utils.MakeArg(args, "[--field=<field>]", l.Field)
    args = utils.MakeArg(args, "[--fields=<fields>]", l.Fields)
    args = utils.MakeArg(args, "[--format=<format>]", l.Format)
    return args
}

