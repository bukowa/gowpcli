/*
## OPTIONS
	<taxonomy>
	: Taxonomy of the term to get
	<term>
	: ID or slug of the term to get
	[--by=<field>]
	: Explicitly handle the term value as a slug or id.
	---
	default: id
	options:
	  - slug
	  - id
	---
	[--field=<field>]
	: Instead of returning the whole term, returns the value of a single field.
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
	    # Get details about a category with id 199.
	    $ wp term get category 199 --format=json
	    {"term_id":199,"name":"Apple","slug":"apple","term_group":0,"term_taxonomy_id":199,"taxonomy":"category","description":"A type of fruit","parent":0,"count":0,"filter":"raw"}
	    # Get details about a category with slug apple.
	    $ wp term get category apple --by=slug --format=json
	    {"term_id":199,"name":"Apple","slug":"apple","term_group":0,"term_taxonomy_id":199,"taxonomy":"category","description":"A type of fruit","parent":0,"count":0,"filter":"raw"}
	
 */
package get
import utils "github.com/bukowa/gowpcli"

// Get //Gets details about a term.
type Get struct {
    Taxonomy string // <taxonomy>
    Term string // <term>
    By string // [--by=<field>]
    Field string // [--field=<field>]
    Fields string // [--fields=<fields>]
    Format string // [--format=<format>]
}

func (g Get) Args() []string {
    var args = []string{"term", "get"}
    args = utils.MakeArg(args, "<taxonomy>", g.Taxonomy)
    args = utils.MakeArg(args, "<term>", g.Term)
    args = utils.MakeArg(args, "[--by=<field>]", g.By)
    args = utils.MakeArg(args, "[--field=<field>]", g.Field)
    args = utils.MakeArg(args, "[--fields=<fields>]", g.Fields)
    args = utils.MakeArg(args, "[--format=<format>]", g.Format)
    return args
}

