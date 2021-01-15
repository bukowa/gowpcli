/*
## OPTIONS
	[--fields=<fields>]
	: Limit the output to specific object fields.
	[--field=<field>]
	: Prints the value of a single field.
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
## AVAILABLE FIELDS
	These fields will be displayed by default for each role:
	* name
	* role
	There are no optional fields.
## EXAMPLES
	    # List roles.
	    $ wp role list --fields=role --format=csv
	    role
	    administrator
	    editor
	    author
	    contributor
	    subscriber
	
 */
package list
import utils "github.com/bukowa/gowpcli"

// List //Lists all roles.
type List struct {
    Fields string // [--fields=<fields>]
    Field string // [--field=<field>]
    Format string // [--format=<format>]
}

func (l List) Args() []string {
    var args = []string{"role", "list"}
    args = utils.MakeArg(args, "[--fields=<fields>]", l.Fields)
    args = utils.MakeArg(args, "[--field=<field>]", l.Field)
    args = utils.MakeArg(args, "[--format=<format>]", l.Format)
    return args
}

