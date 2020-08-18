/*
## OPTIONS
	[--field=<field>]
	: Display the value of a single field
	[--<field>=<value>]
	: Filter results by key=value pairs.
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
	---
## AVAILABLE FIELDS
	These fields will be displayed by default for each translation:
	* language
	* english_name
	* native_name
	* status
	* update
	* updated
## EXAMPLES
	    # List language,english_name,status fields of available languages.
	    $ wp language core list --fields=language,english_name,status
	    +----------------+-------------------------+-------------+
	    | language       | english_name            | status      |
	    +----------------+-------------------------+-------------+
	    | ar             | Arabic                  | uninstalled |
	    | ary            | Moroccan Arabic         | uninstalled |
	    | az             | Azerbaijani             | uninstalled |
	
 */
package list
import utils "github.com/bukowa/gowpcli"

// List //Lists all available languages.
type List struct {
    Field string // [--field=<field>]
    FieldMap map[string]string // [--<field>=<value>]
    Fields string // [--fields=<fields>]
    Format string // [--format=<format>]
}

func (l List) Args() []string {
    var args = []string{"language", "core", "list"}
    args = utils.MakeArg(args, "[--field=<field>]", l.Field)
    args = utils.MakeArg(args, "[--<field>=<value>]", l.FieldMap)
    args = utils.MakeArg(args, "[--fields=<fields>]", l.Fields)
    args = utils.MakeArg(args, "[--format=<format>]", l.Format)
    return args
}

