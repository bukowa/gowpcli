/*
## OPTIONS
	<id>
	: The ID of the object.
	<taxonomy>
	: The name of the term's taxonomy.
	[<term>...]
	: The name of the term or terms to be removed from the object.
	[--by=<field>]
	: Explicitly handle the term value as a slug or id.
	---
	options:
	  - slug
	  - id
	---
	[--all]
	: Remove all terms from the object.
	
 */
package remove
import utils "github.com/bukowa/gowpcli"

// Remove //Remove a term from an object.
type Remove struct {
    Id string // <id>
    Taxonomy string // <taxonomy>
    Term []string // [<term>...]
    By string // [--by=<field>]
    All bool // [--all]
}

func (r Remove) Args() []string {
    var args = []string{"post", "term", "remove"}
    args = utils.MakeArg(args, "<id>", r.Id)
    args = utils.MakeArg(args, "<taxonomy>", r.Taxonomy)
    args = utils.MakeArg(args, "[<term>...]", r.Term)
    args = utils.MakeArg(args, "[--by=<field>]", r.By)
    args = utils.MakeArg(args, "[--all]", r.All)
    return args
}

