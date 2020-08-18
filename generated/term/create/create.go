/*
## OPTIONS
	<taxonomy>
	: Taxonomy for the new term.
	<term>
	: A name for the new term.
	[--slug=<slug>]
	: A unique slug for the new term. Defaults to sanitized version of name.
	[--description=<description>]
	: A description for the new term.
	[--parent=<term-id>]
	: A parent for the new term.
	[--porcelain]
	: Output just the new term id.
## EXAMPLES
	    # Create a new category "Apple" with a description.
	    $ wp term create category Apple --description="A type of fruit"
	    Success: Created category 199.
	
 */
package create
import utils "github.com/bukowa/gowpcli"

// Create //Creates a new term.
type Create struct {
    Taxonomy string // <taxonomy>
    Term string // <term>
    Slug string // [--slug=<slug>]
    Description string // [--description=<description>]
    Parent string // [--parent=<term-id>]
    Porcelain bool // [--porcelain]
}

func (c Create) Args() []string {
    var args = []string{"term", "create"}
    args = utils.MakeArg(args, "<taxonomy>", c.Taxonomy)
    args = utils.MakeArg(args, "<term>", c.Term)
    args = utils.MakeArg(args, "[--slug=<slug>]", c.Slug)
    args = utils.MakeArg(args, "[--description=<description>]", c.Description)
    args = utils.MakeArg(args, "[--parent=<term-id>]", c.Parent)
    args = utils.MakeArg(args, "[--porcelain]", c.Porcelain)
    return args
}

