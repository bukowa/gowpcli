/*
## OPTIONS
	[--fields=<fields>]
	: Limit the output to specific object fields.
	[--format=<format>]
	: Render output in a particular format.
	---
	default: table
	options:
	  - table
	  - csv
	  - json
	  - count
	  - ids
	  - yaml
	---
## AVAILABLE FIELDS
	These fields will be displayed by default for each menu:
	* term_id
	* name
	* slug
	* count
	These fields are optionally available:
	* term_group
	* term_taxonomy_id
	* taxonomy
	* description
	* parent
	* locations
## EXAMPLES
	    $ wp menu list
	    +---------+----------+----------+-----------+-------+
	    | term_id | name     | slug     | locations | count |
	    +---------+----------+----------+-----------+-------+
	    | 200     | My Menu  | my-menu  |           | 0     |
	    | 177     | Top Menu | top-menu | primary   | 7     |
	    +---------+----------+----------+-----------+-------+
	
 */
package list
import utils "github.com/bukowa/gowpcli"

// List //Gets a list of menus.
type List struct {
    Fields string // [--fields=<fields>]
    Format string // [--format=<format>]
}

func (l List) Args() []string {
    var args = []string{"menu", "list"}
    args = utils.MakeArg(args, "[--fields=<fields>]", l.Fields)
    args = utils.MakeArg(args, "[--format=<format>]", l.Format)
    return args
}

