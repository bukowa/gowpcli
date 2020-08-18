/*
## OPTIONS
	[<site-id>]
	: The id of the site to delete. If not provided, you must set the --slug parameter.
	[--slug=<slug>]
	: Path of the blog to be deleted. Subdomain on subdomain installs, directory on subdirectory installs.
	[--yes]
	: Answer yes to the confirmation message.
	[--keep-tables]
	: Delete the blog from the list, but don't drop it's tables.
## EXAMPLES
	    $ wp site delete 123
	    Are you sure you want to delete the http://www.example.com/example site? [y/n] y
	    Success: The site at 'http://www.example.com/example' was deleted.
	
 */
package delete
import utils "github.com/bukowa/gowpcli"

// Delete //Deletes a site in a multisite installation.
type Delete struct {
    SiteId string // [<site-id>]
    Slug string // [--slug=<slug>]
    Yes bool // [--yes]
    KeepTables bool // [--keep-tables]
}

func (d Delete) Args() []string {
    var args = []string{"site", "delete"}
    args = utils.MakeArg(args, "[<site-id>]", d.SiteId)
    args = utils.MakeArg(args, "[--slug=<slug>]", d.Slug)
    args = utils.MakeArg(args, "[--yes]", d.Yes)
    args = utils.MakeArg(args, "[--keep-tables]", d.KeepTables)
    return args
}

