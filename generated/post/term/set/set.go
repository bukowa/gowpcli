/*
## INFO
	Replaces existing terms on the object.
	<id>
	: The ID of the object.
	<taxonomy>
	: The name of the taxonomy type to be updated.
	<term>...
	: The slug of the term or terms to be updated.
	[--by=<field>]
	: Explicitly handle the term value as a slug or id.
	---
	options:
	  - slug
	  - id
	---
	
 */
package set
import utils "github.com/bukowa/gowpcli"

// Set //Set object terms.
type Set struct {
    Id string // <id>
    Taxonomy string // <taxonomy>
    Term []string // <term>...
    By string // [--by=<field>]
}

func (s Set) Args() []string {
    var args = []string{"post", "term", "set"}
    args = utils.MakeArg(args, "<id>", s.Id)
    args = utils.MakeArg(args, "<taxonomy>", s.Taxonomy)
    args = utils.MakeArg(args, "<term>...", s.Term)
    args = utils.MakeArg(args, "[--by=<field>]", s.By)
    return args
}

