/*
## OPTIONS
	[--field=<field>]
	: Display the value of a single field
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
	These fields will be displayed by default for each handler:
	* id
	* regex
	These fields are optionally available:
	* callback
	* priority
## EXAMPLES
	    # List id,regex,priority fields of available handlers.
	    $ wp embed handler list --fields=priority,id
	    +----------+-------------------+
	    | priority | id                |
	    +----------+-------------------+
	    | 10       | youtube_embed_url |
	    | 9999     | audio             |
	    | 9999     | video             |
	
 */
package list
import utils "github.com/bukowa/gowpcli"

// List //Lists all available embed handlers.
type List struct {
    Field string // [--field=<field>]
    Fields string // [--fields=<fields>]
    Format string // [--format=<format>]
}

func (l List) Args() []string {
    var args = []string{"embed", "handler", "list"}
    args = utils.MakeArg(args, "[--field=<field>]", l.Field)
    args = utils.MakeArg(args, "[--fields=<fields>]", l.Fields)
    args = utils.MakeArg(args, "[--format=<format>]", l.Format)
    return args
}

