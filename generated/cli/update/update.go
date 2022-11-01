/*
## INFO
	Default behavior is to check the releases API for the newest stable
	version, and prompt if one is available.
	Use `--stable` to install or reinstall the latest stable version.
	Use `--nightly` to install the latest built version of the master branch.
	While not recommended for production, nightly contains the latest and
	greatest, and should be stable enough for development and staging
	environments.
	Only works for the Phar installation mechanism.
## OPTIONS
	[--patch]
	: Only perform patch updates.
	[--minor]
	: Only perform minor updates.
	[--major]
	: Only perform major updates.
	[--stable]
	: Update to the latest stable release. Skips update check.
	[--nightly]
	: Update to the latest built version of the master branch. Potentially unstable.
	[--yes]
	: Do not prompt for confirmation.
	[--insecure]
	: Retry without certificate validation if TLS handshake fails. Note: This makes the request vulnerable to a MITM attack.
## EXAMPLES
	    # Update CLI.
	    $ wp cli update
	    You have version 0.24.0. Would you like to update to 0.24.1? [y/n] y
	    Downloading from https://github.com/wp-cli/wp-cli/releases/download/v0.24.1/wp-cli-0.24.1.phar...
	    New version works. Proceeding to replace.
	    Success: Updated WP-CLI to 0.24.1.
	
 */
package update
import utils "github.com/bukowa/gowpcli"

// Update //Updates WP-CLI to the latest release.
type Update struct {
    Patch bool // [--patch]
    Minor bool // [--minor]
    Major bool // [--major]
    Stable bool // [--stable]
    Nightly bool // [--nightly]
    Yes bool // [--yes]
    Insecure bool // [--insecure]
}

func (u Update) Args() []string {
    var args = []string{"cli", "update"}
    args = utils.MakeArg(args, "[--patch]", u.Patch)
    args = utils.MakeArg(args, "[--minor]", u.Minor)
    args = utils.MakeArg(args, "[--major]", u.Major)
    args = utils.MakeArg(args, "[--stable]", u.Stable)
    args = utils.MakeArg(args, "[--nightly]", u.Nightly)
    args = utils.MakeArg(args, "[--yes]", u.Yes)
    args = utils.MakeArg(args, "[--insecure]", u.Insecure)
    return args
}

