/*
## OPTIONS
	<user>
	: The user login, user email, or user ID of the user to get application passwords for.
	[--<field>=<value>]
	: Filter the list by a specific field.
	[--field=<field>]
	: Prints the value of a single field for each application password.
	[--fields=<fields>]
	: Limit the output to specific fields.
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
	default: created
	options:
	 - uuid
	 - app_id
	 - name
	 - password
	 - created
	 - last_used
	 - last_ip
	---
	[--order=<order>]
	: Set ascending or descending order.
	---
	default: desc
	options:
	 - asc
	 - desc
	---
## EXAMPLES
	    # List user application passwords and only show app name and password hash
	    $ wp user application-password list 123 --fields=name,password
	    +--------+------------------------------------+
	    | name   | password                           |
	    +--------+------------------------------------+
	    | myapp  | $P$BVGeou1CUot114YohIemgpwxQCzb8O/ |
	    +--------+------------------------------------+
	
 */
package list
import utils "github.com/bukowa/gowpcli"

// List //Lists all application passwords associated with a user.
type List struct {
    User string // <user>
    FieldMap map[string]string // [--<field>=<value>]
    Field string // [--field=<field>]
    Fields string // [--fields=<fields>]
    Format string // [--format=<format>]
    Orderby string // [--orderby=<fields>]
    Order string // [--order=<order>]
}

func (l List) Args() []string {
    var args = []string{"user", "application-password", "list"}
    args = utils.MakeArg(args, "<user>", l.User)
    args = utils.MakeArg(args, "[--<field>=<value>]", l.FieldMap)
    args = utils.MakeArg(args, "[--field=<field>]", l.Field)
    args = utils.MakeArg(args, "[--fields=<fields>]", l.Fields)
    args = utils.MakeArg(args, "[--format=<format>]", l.Format)
    args = utils.MakeArg(args, "[--orderby=<fields>]", l.Orderby)
    args = utils.MakeArg(args, "[--order=<order>]", l.Order)
    return args
}

