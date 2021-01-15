/*
## OPTIONS
	<user>
	: The user login, user email, or user ID of the user to get metadata for.
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
	  - count
	  - yaml
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
## EXAMPLES
	    # List user meta
	    $ wp user meta list 123 --keys=nickname,description,wp_capabilities
	    +---------+-----------------+--------------------------------+
	    | user_id | meta_key        | meta_value                     |
	    +---------+-----------------+--------------------------------+
	    | 123     | nickname        | supervisor                     |
	    | 123     | description     | Mary is a WordPress developer. |
	    | 123     | wp_capabilities | {"administrator":true}         |
	    +---------+-----------------+--------------------------------+
	
 */
package list
import utils "github.com/bukowa/gowpcli"

// List //Lists all metadata associated with a user.
type List struct {
    User string // <user>
    Keys string // [--keys=<keys>]
    Fields string // [--fields=<fields>]
    Format string // [--format=<format>]
    Orderby string // [--orderby=<fields>]
    Order string // [--order=<order>]
    Unserialize bool // [--unserialize]
}

func (l List) Args() []string {
    var args = []string{"user", "meta", "list"}
    args = utils.MakeArg(args, "<user>", l.User)
    args = utils.MakeArg(args, "[--keys=<keys>]", l.Keys)
    args = utils.MakeArg(args, "[--fields=<fields>]", l.Fields)
    args = utils.MakeArg(args, "[--format=<format>]", l.Format)
    args = utils.MakeArg(args, "[--orderby=<fields>]", l.Orderby)
    args = utils.MakeArg(args, "[--order=<order>]", l.Order)
    args = utils.MakeArg(args, "[--unserialize]", l.Unserialize)
    return args
}

