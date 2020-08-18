/*
## OPTIONS
	<id>
	: The ID of the post to get.
	[--field=<field>]
	: Instead of returning the whole post, returns the value of a single field.
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
## EXAMPLES
	    # Save the post content to a file
	    $ wp post get 123 --field=content > file.txt
	
 */
package get
import utils "github.com/bukowa/gowpcli"

// Get //Gets details about a post.
type Get struct {
    Id string // <id>
    Field string // [--field=<field>]
    Fields string // [--fields=<fields>]
    Format string // [--format=<format>]
}

func (g Get) Args() []string {
    var args = []string{"post", "get"}
    args = utils.MakeArg(args, "<id>", g.Id)
    args = utils.MakeArg(args, "[--field=<field>]", g.Field)
    args = utils.MakeArg(args, "[--fields=<fields>]", g.Fields)
    args = utils.MakeArg(args, "[--format=<format>]", g.Format)
    return args
}

