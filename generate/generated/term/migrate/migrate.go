/*
## OPTIONS
	<term>
	: Slug or ID of the term to migrate.
	[--by=<field>]
	: Explicitly handle the term value as a slug or id.
	---
	default: id
	options:
	  - slug
	  - id
	---
	[--from=<taxonomy>]
	: Taxonomy slug of the term to migrate.
	[--to=<taxonomy>]
	: Taxonomy slug to migrate to.
## EXAMPLES
	    # Migrate a category's term (video) to tag taxonomy.
	    $ wp term migrate 9190 --from=category --to=post_tag
	    Term '9190' migrated!
	    Old instance of term '9190' removed from its original taxonomy.
	    Success: Migrated the term '9190' from taxonomy 'category' to taxonomy 'post_tag' for 1 posts
	
 */
package migrate
import utils "github.com/bukowa/gowpcli"

// Migrate //Migrate a term of a taxonomy to another taxonomy.
type Migrate struct {
    Term string // <term>
    By string // [--by=<field>]
    From string // [--from=<taxonomy>]
    To string // [--to=<taxonomy>]
}

func (m Migrate) Args() []string {
    var args = []string{"term", "migrate"}
    args = utils.MakeArg(args, "<term>", m.Term)
    args = utils.MakeArg(args, "[--by=<field>]", m.By)
    args = utils.MakeArg(args, "[--from=<taxonomy>]", m.From)
    args = utils.MakeArg(args, "[--to=<taxonomy>]", m.To)
    return args
}

