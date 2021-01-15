/*
## OPTIONS
	<post-type>
	: Post type slug
	[--field=<field>]
	: Instead of returning the whole taxonomy, returns the value of a single field.
	[--fields=<fields>]
	: Limit the output to specific fields. Defaults to all fields.
	[--format=<format>]
	: Render output in a particular format.
	---
	default: table
	options:
	  - table
	  - csv
	  - json
	  - yaml
	---
## AVAILABLE FIELDS
	These fields will be displayed by default for the specified post type:
	* name
	* label
	* description
	* hierarchical
	* public
	* capability_type
	* labels
	* cap
	* supports
	These fields are optionally available:
	* count
## EXAMPLES
	    # Get details about the 'page' post type.
	    $ wp post-type get page --fields=name,label,hierarchical --format=json
	    {"name":"page","label":"Pages","hierarchical":true}
	
 */
package get
import utils "github.com/bukowa/gowpcli"

// Get //Gets details about a registered post type.
type Get struct {
    PostType string // <post-type>
    Field string // [--field=<field>]
    Fields string // [--fields=<fields>]
    Format string // [--format=<format>]
}

func (g Get) Args() []string {
    var args = []string{"post-type", "get"}
    args = utils.MakeArg(args, "<post-type>", g.PostType)
    args = utils.MakeArg(args, "[--field=<field>]", g.Field)
    args = utils.MakeArg(args, "[--fields=<fields>]", g.Fields)
    args = utils.MakeArg(args, "[--format=<format>]", g.Format)
    return args
}

