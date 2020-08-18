/*
## OPTIONS
	<sidebar-id>
	: ID for the corresponding sidebar.
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
	These fields will be displayed by default for each widget:
	* name
	* id
	* position
	* options
	There are no optionally available fields.
## EXAMPLES
	    $ wp widget list sidebar-1 --fields=name,id --format=csv
	    name,id
	    meta,meta-5
	    search,search-3
	
 */
package list
import utils "github.com/bukowa/gowpcli"

// List //Lists widgets associated with a sidebar.
type List struct {
    SidebarId string // <sidebar-id>
    Fields string // [--fields=<fields>]
    Format string // [--format=<format>]
}

func (l List) Args() []string {
    var args = []string{"widget", "list"}
    args = utils.MakeArg(args, "<sidebar-id>", l.SidebarId)
    args = utils.MakeArg(args, "[--fields=<fields>]", l.Fields)
    args = utils.MakeArg(args, "[--format=<format>]", l.Format)
    return args
}

