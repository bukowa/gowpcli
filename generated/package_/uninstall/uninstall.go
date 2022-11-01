/*
## OPTIONS
	<name>
	: Name of the package to uninstall.
	[--insecure]
	: Retry downloads without certificate validation if TLS handshake fails. Note: This makes the request vulnerable to a MITM attack.
## EXAMPLES
	    $ wp package uninstall wp-cli/server-command
	    Removing require statement from /home/person/.wp-cli/packages/composer.json
	    Deleting package directory /home/person/.wp-cli/packages/vendor/wp-cli/server-command
	    Regenerating Composer autoload.
	    Success: Uninstalled package.
	
 */
package uninstall
import utils "github.com/bukowa/gowpcli"

// Uninstall //Uninstalls a WP-CLI package.
type Uninstall struct {
    Name string // <name>
    Insecure bool // [--insecure]
}

func (u Uninstall) Args() []string {
    var args = []string{"package", "uninstall"}
    args = utils.MakeArg(args, "<name>", u.Name)
    args = utils.MakeArg(args, "[--insecure]", u.Insecure)
    return args
}

