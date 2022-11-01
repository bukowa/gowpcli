/*
## OPTIONS
	[<plugin>...]
	: One or more plugins to verify.
	[--all]
	: If set, all plugins will be verified.
	[--strict]
	: If set, even "soft changes" like readme.txt changes will trigger
	checksum errors.
	[--format=<format>]
	: Render output in a specific format.
	---
	default: table
	options:
	  - table
	  - json
	  - csv
	  - yaml
	  - count
	---
	[--insecure]
	: Retry downloads without certificate validation if TLS handshake fails. Note: This makes the request vulnerable to a MITM attack.
## EXAMPLES
	    # Verify the checksums of all installed plugins
	    $ wp plugin verify-checksums --all
	    Success: Verified 8 of 8 plugins.
	    # Verify the checksums of a single plugin, Akismet in this case
	    $ wp plugin verify-checksums akismet
	    Success: Verified 1 of 1 plugins.
	
 */
package verifychecksums
import utils "github.com/bukowa/gowpcli"

// VerifyChecksums //Verifies plugin files against WordPress.org's checksums.
type VerifyChecksums struct {
    Plugin []string // [<plugin>...]
    All bool // [--all]
    Strict bool // [--strict]
    Format string // [--format=<format>]
    Insecure bool // [--insecure]
}

func (v VerifyChecksums) Args() []string {
    var args = []string{"plugin", "verify-checksums"}
    args = utils.MakeArg(args, "[<plugin>...]", v.Plugin)
    args = utils.MakeArg(args, "[--all]", v.All)
    args = utils.MakeArg(args, "[--strict]", v.Strict)
    args = utils.MakeArg(args, "[--format=<format>]", v.Format)
    args = utils.MakeArg(args, "[--insecure]", v.Insecure)
    return args
}

