/*
## OPTIONS
	[<theme>...]
	: One or more themes to list languages for.
	[--all]
	: If set, available languages for all themes will be listed.
	[--field=<field>]
	: Display the value of a single field.
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
	* theme
	* language
	* english_name
	* native_name
	* status
	* update
	* updated
## EXAMPLES
	    # List language,english_name,status fields of available languages.
	    $ wp language theme list --fields=language,english_name,status
	    +----------------+-------------------------+-------------+
	    | language       | english_name            | status      |
	    +----------------+-------------------------+-------------+
	    | ar             | Arabic                  | uninstalled |
	    | ary            | Moroccan Arabic         | uninstalled |
	    | az             | Azerbaijani             | uninstalled |
	
 */
package list
import utils "github.com/bukowa/gowpcli"

// List //Lists all available languages for one or more themes.
type List struct {
    Theme []string // [<theme>...]
    All bool // [--all]
    Field string // [--field=<field>]
    FieldMap map[string]string // [--<field>=<value>]
    Fields string // [--fields=<fields>]
    Format string // [--format=<format>]
}

func (l List) Args() []string {
    var args = []string{"language", "theme", "list"}
    args = utils.MakeArg(args, "[<theme>...]", l.Theme)
    args = utils.MakeArg(args, "[--all]", l.All)
    args = utils.MakeArg(args, "[--field=<field>]", l.Field)
    args = utils.MakeArg(args, "[--<field>=<value>]", l.FieldMap)
    args = utils.MakeArg(args, "[--fields=<fields>]", l.Fields)
    args = utils.MakeArg(args, "[--format=<format>]", l.Format)
    return args
}

