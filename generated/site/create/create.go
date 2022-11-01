/*
## OPTIONS
	--slug=<slug>
	: Path for the new site. Subdomain on subdomain installs, directory on subdirectory installs.
	[--title=<title>]
	: Title of the new site. Default: prettified slug.
	[--email=<email>]
	: Email for admin user. User will be created if none exists. Assignment to super admin if not included.
	[--network_id=<network-id>]
	: Network to associate new site with. Defaults to current network (typically 1).
	[--private]
	: If set, the new site will be non-public (not indexed)
	[--porcelain]
	: If set, only the site id will be output on success.
## EXAMPLES
	    $ wp site create --slug=example
	    Success: Site 3 created: http://www.example.com/example/
	
 */
package create
import utils "github.com/bukowa/gowpcli"

// Create //Creates a site in a multisite installation.
type Create struct {
    Slug string // --slug=<slug>
    Title string // [--title=<title>]
    Email string // [--email=<email>]
    NetworkId string // [--network_id=<network-id>]
    Private bool // [--private]
    Porcelain bool // [--porcelain]
}

func (c Create) Args() []string {
    var args = []string{"site", "create"}
    args = utils.MakeArg(args, "--slug=<slug>", c.Slug)
    args = utils.MakeArg(args, "[--title=<title>]", c.Title)
    args = utils.MakeArg(args, "[--email=<email>]", c.Email)
    args = utils.MakeArg(args, "[--network_id=<network-id>]", c.NetworkId)
    args = utils.MakeArg(args, "[--private]", c.Private)
    args = utils.MakeArg(args, "[--porcelain]", c.Porcelain)
    return args
}

