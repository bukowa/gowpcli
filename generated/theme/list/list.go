/*
## OPTIONS
	[--<field>=<value>]
	: Filter results based on the value of a field.
	[--field=<field>]
	: Prints the value of a single field for each theme.
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
	  - yaml
	---
	[--status=<status>]
	: Filter the output by theme status.
	---
	options:
	  - active
	  - parent
	  - inactive
	---
	[--skip-update-check]
	: If set, the theme update check will be skipped.
## AVAILABLE FIELDS
	These fields will be displayed by default for each theme:
	* name
	* status
	* update
	* version
	These fields are optionally available:
	* update_version
	* update_package
	* update_id
	* title
	* description
## EXAMPLES
	    # List themes
	    $ wp theme list --status=inactive --format=csv
	    name,status,update,version
	    twentyfourteen,inactive,none,1.7
	    twentysixteen,inactive,available,1.1
	
 */
package list
import utils "github.com/bukowa/gowpcli"

// List //Gets a list of themes.
type List struct {
    FieldMap map[string]string // [--<field>=<value>]
    Field string // [--field=<field>]
    Fields string // [--fields=<fields>]
    Format string // [--format=<format>]
    Status string // [--status=<status>]
    SkipUpdateCheck bool // [--skip-update-check]
}

func (l List) Args() []string {
    var args = []string{"theme", "list"}
    args = utils.MakeArg(args, "[--<field>=<value>]", l.FieldMap)
    args = utils.MakeArg(args, "[--field=<field>]", l.Field)
    args = utils.MakeArg(args, "[--fields=<fields>]", l.Fields)
    args = utils.MakeArg(args, "[--format=<format>]", l.Format)
    args = utils.MakeArg(args, "[--status=<status>]", l.Status)
    args = utils.MakeArg(args, "[--skip-update-check]", l.SkipUpdateCheck)
    return args
}

