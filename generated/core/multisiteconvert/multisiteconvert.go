/*
## INFO
	Creates the multisite database tables, and adds the multisite constants
	to wp-config.php.
	For those using WordPress with Apache, remember to update the `.htaccess`
	file with the appropriate multisite rewrite rules.
	[Review the multisite documentation](https://wordpress.org/support/article/create-a-network/)
	for more details about how multisite works.
## OPTIONS
	[--title=<network-title>]
	: The title of the new network.
	[--base=<url-path>]
	: Base path after the domain name that each site url will start with.
	---
	default: /
	---
	[--subdomains]
	: If passed, the network will use subdomains, instead of subdirectories. Doesn't work with 'localhost'.
## EXAMPLES
	    $ wp core multisite-convert
	    Set up multisite database tables.
	    Added multisite constants to wp-config.php.
	    Success: Network installed. Don't forget to set up rewrite rules.
	
 */
package multisiteconvert
import utils "github.com/bukowa/gowpcli"

// MultisiteConvert //Transforms an existing single-site installation into a multisite installation.
type MultisiteConvert struct {
    Title string // [--title=<network-title>]
    Base string // [--base=<url-path>]
    Subdomains bool // [--subdomains]
}

func (m MultisiteConvert) Args() []string {
    var args = []string{"core", "multisite-convert"}
    args = utils.MakeArg(args, "[--title=<network-title>]", m.Title)
    args = utils.MakeArg(args, "[--base=<url-path>]", m.Base)
    args = utils.MakeArg(args, "[--subdomains]", m.Subdomains)
    return args
}

