/*
## OPTIONS
	<taxonomy>
	: Taxonomy of the term to update.
	<term>
	: ID or slug for the term to update.
	[--by=<field>]
	: Explicitly handle the term value as a slug or id.
	---
	default: id
	options:
	  - slug
	  - id
	---
	[--name=<name>]
	: A new name for the term.
	[--slug=<slug>]
	: A new slug for the term.
	[--description=<description>]
	: A new description for the term.
	[--parent=<term-id>]
	: A new parent for the term.
## EXAMPLES
	    # Change category with id 15 to use the name "Apple"
	    $ wp term update category 15 --name=Apple
	    Success: Term updated.
	    # Change category with slug apple to use the name "Apple"
	    $ wp term update category apple --by=slug --name=Apple
	    Success: Term updated.
	
 */
package update
import utils "github.com/bukowa/gowpcli"

// Update //Updates an existing term.
type Update struct {
    Taxonomy string // <taxonomy>
    Term string // <term>
    By string // [--by=<field>]
    Name string // [--name=<name>]
    Slug string // [--slug=<slug>]
    Description string // [--description=<description>]
    Parent string // [--parent=<term-id>]
}

func (u Update) Args() []string {
    var args = []string{"term", "update"}
    args = utils.MakeArg(args, "<taxonomy>", u.Taxonomy)
    args = utils.MakeArg(args, "<term>", u.Term)
    args = utils.MakeArg(args, "[--by=<field>]", u.By)
    args = utils.MakeArg(args, "[--name=<name>]", u.Name)
    args = utils.MakeArg(args, "[--slug=<slug>]", u.Slug)
    args = utils.MakeArg(args, "[--description=<description>]", u.Description)
    args = utils.MakeArg(args, "[--parent=<term-id>]", u.Parent)
    return args
}

