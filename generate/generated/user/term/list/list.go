/*
## INFO
	<id>
	: ID for the object.
	<taxonomy>...
	: One or more taxonomies to list.
	[--field=<field>]
	: Prints the value of a single field for each term.
	[--fields=<fields>]
	: Limit the output to specific row fields.
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
	  - ids
	---
## AVAILABLE FIELDS
	These fields will be displayed by default for each term:
	* term_id
	* name
	* slug
	* taxonomy
	These fields are optionally available:
	* term_taxonomy_id
	* description
	* term_group
	* parent
	* count
	
 */
package list
import utils "github.com/bukowa/gowpcli"

// List //List all terms associated with an object.
type List struct {
    Id string // <id>
    Taxonomy []string // <taxonomy>...
    Field string // [--field=<field>]
    Fields string // [--fields=<fields>]
    Format string // [--format=<format>]
}

func (l List) Args() []string {
    var args = []string{"user", "term", "list"}
    args = utils.MakeArg(args, "<id>", l.Id)
    args = utils.MakeArg(args, "<taxonomy>...", l.Taxonomy)
    args = utils.MakeArg(args, "[--field=<field>]", l.Field)
    args = utils.MakeArg(args, "[--fields=<fields>]", l.Fields)
    args = utils.MakeArg(args, "[--format=<format>]", l.Format)
    return args
}

