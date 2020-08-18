/*
## OPTIONS
	<id>
	: ID for the object.
	[--keys=<keys>]
	: Limit output to metadata of specific keys.
	[--fields=<fields>]
	: Limit the output to specific row fields. Defaults to id,meta_key,meta_value.
	[--format=<format>]
	: Render output in a particular format.
	---
	default: table
	options:
	  - table
	  - csv
	  - json
	  - yaml
	  - count
	---
	[--orderby=<fields>]
	: Set orderby which field.
	---
	default: id
	options:
	 - id
	 - meta_key
	 - meta_value
	---
	[--order=<order>]
	: Set ascending or descending order.
	---
	default: asc
	options:
	 - asc
	 - desc
	---
	[--unserialize]
	: Unserialize meta_value output.
	
 */
package list
import utils "github.com/bukowa/gowpcli"

// List //List all metadata associated with an object.
type List struct {
    Id string // <id>
    Keys string // [--keys=<keys>]
    Fields string // [--fields=<fields>]
    Format string // [--format=<format>]
    Orderby string // [--orderby=<fields>]
    Order string // [--order=<order>]
    Unserialize bool // [--unserialize]
}

func (l List) Args() []string {
    var args = []string{"comment", "meta", "list"}
    args = utils.MakeArg(args, "<id>", l.Id)
    args = utils.MakeArg(args, "[--keys=<keys>]", l.Keys)
    args = utils.MakeArg(args, "[--fields=<fields>]", l.Fields)
    args = utils.MakeArg(args, "[--format=<format>]", l.Format)
    args = utils.MakeArg(args, "[--orderby=<fields>]", l.Orderby)
    args = utils.MakeArg(args, "[--order=<order>]", l.Order)
    args = utils.MakeArg(args, "[--unserialize]", l.Unserialize)
    return args
}

