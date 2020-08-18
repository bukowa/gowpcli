/*
## OPTIONS
	[--search=<pattern>]
	: Use wildcards ( * and ? ) to match option name.
	[--site_id=<id>]
	: Limit options to those of a particular site id.
	[--field=<field>]
	: Prints the value of a single field.
	[--fields=<fields>]
	: Limit the output to specific object fields.
	[--format=<format>]
	: The serialization format for the value. total_bytes displays the total size of matching options in bytes.
	---
	default: table
	options:
	  - table
	  - json
	  - csv
	  - count
	  - yaml
	  - total_bytes
	---
## AVAILABLE FIELDS
	This field will be displayed by default for each matching option:
	* meta_key
	* meta_value
	These fields are optionally available:
	* meta_id
	* site_id
	* size_bytes
## EXAMPLES
	    # List all site options begining with "i2f_"
	    $ wp site option list --search="i2f_*"
	    +-------------+--------------+
	    | meta_key    | meta_value   |
	    +-------------+--------------+
	    | i2f_version | 0.1.0        |
	    +-------------+--------------+
	
 */
package list
import utils "github.com/bukowa/gowpcli"

// List //Lists site options.
type List struct {
    Search string // [--search=<pattern>]
    SiteId string // [--site_id=<id>]
    Field string // [--field=<field>]
    Fields string // [--fields=<fields>]
    Format string // [--format=<format>]
}

func (l List) Args() []string {
    var args = []string{"site", "option", "list"}
    args = utils.MakeArg(args, "[--search=<pattern>]", l.Search)
    args = utils.MakeArg(args, "[--site_id=<id>]", l.SiteId)
    args = utils.MakeArg(args, "[--field=<field>]", l.Field)
    args = utils.MakeArg(args, "[--fields=<fields>]", l.Fields)
    args = utils.MakeArg(args, "[--format=<format>]", l.Format)
    return args
}

