/*
## INFO
	Append the term to the existing set of terms on the object.
	<id>
	: The ID of the object.
	<taxonomy>
	: The name of the taxonomy type to be added.
	<term>...
	: The slug of the term or terms to be added.
	[--by=<field>]
	: Explicitly handle the term value as a slug or id.
	---
	options:
	  - slug
	  - id
	---
	
 */
package add
import utils "github.com/bukowa/gowpcli"

// Add //Add a term to an object.
type Add struct {
    Id string // <id>
    Taxonomy string // <taxonomy>
    Term []string // <term>...
    By string // [--by=<field>]
}

func (a Add) Args() []string {
    var args = []string{"user", "term", "add"}
    args = utils.MakeArg(args, "<id>", a.Id)
    args = utils.MakeArg(args, "<taxonomy>", a.Taxonomy)
    args = utils.MakeArg(args, "<term>...", a.Term)
    args = utils.MakeArg(args, "[--by=<field>]", a.By)
    return args
}

