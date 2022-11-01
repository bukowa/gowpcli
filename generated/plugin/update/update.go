/*
## OPTIONS
	[<plugin>...]
	: One or more plugins to update.
	[--all]
	: If set, all plugins that have updates will be updated.
	[--exclude=<name>]
	: Comma separated list of plugin names that should be excluded from updating.
	[--minor]
	: Only perform updates for minor releases (e.g. from 1.3 to 1.4 instead of 2.0)
	[--patch]
	: Only perform updates for patch releases (e.g. from 1.3 to 1.3.3 instead of 1.4)
	[--format=<format>]
	: Render output in a particular format.
	---
	default: table
	options:
	  - table
	  - csv
	  - json
	  - summary
	---
	[--version=<version>]
	: If set, the plugin will be updated to the specified version.
	[--dry-run]
	: Preview which plugins would be updated.
	[--insecure]
	: Retry downloads without certificate validation if TLS handshake fails. Note: This makes the request vulnerable to a MITM attack.
## EXAMPLES
	    $ wp plugin update bbpress --version=dev
	    Installing bbPress (Development Version)
	    Downloading install package from https://downloads.wordpress.org/plugin/bbpress.zip...
	    Unpacking the package...
	    Installing the plugin...
	    Removing the old version of the plugin...
	    Plugin updated successfully.
	    Success: Updated 1 of 2 plugins.
	    $ wp plugin update --all
	    Enabling Maintenance mode...
	    Downloading update from https://downloads.wordpress.org/plugin/akismet.3.1.11.zip...
	    Unpacking the update...
	    Installing the latest version...
	    Removing the old version of the plugin...
	    Plugin updated successfully.
	    Downloading update from https://downloads.wordpress.org/plugin/nginx-champuru.3.2.0.zip...
	    Unpacking the update...
	    Installing the latest version...
	    Removing the old version of the plugin...
	    Plugin updated successfully.
	    Disabling Maintenance mode...
	    +------------------------+-------------+-------------+---------+
	    | name                   | old_version | new_version | status  |
	    +------------------------+-------------+-------------+---------+
	    | akismet                | 3.1.3       | 3.1.11      | Updated |
	    | nginx-cache-controller | 3.1.1       | 3.2.0       | Updated |
	    +------------------------+-------------+-------------+---------+
	    Success: Updated 2 of 2 plugins.
	    $ wp plugin update --all --exclude=akismet
	    Enabling Maintenance mode...
	    Downloading update from https://downloads.wordpress.org/plugin/nginx-champuru.3.2.0.zip...
	    Unpacking the update...
	    Installing the latest version...
	    Removing the old version of the plugin...
	    Plugin updated successfully.
	    Disabling Maintenance mode...
	    +------------------------+-------------+-------------+---------+
	    | name                   | old_version | new_version | status  |
	    +------------------------+-------------+-------------+---------+
	    | nginx-cache-controller | 3.1.1       | 3.2.0       | Updated |
	    +------------------------+-------------+-------------+---------+
	
 */
package update
import utils "github.com/bukowa/gowpcli"

// Update //Updates one or more plugins.
type Update struct {
    Plugin []string // [<plugin>...]
    All bool // [--all]
    Exclude string // [--exclude=<name>]
    Minor bool // [--minor]
    Patch bool // [--patch]
    Format string // [--format=<format>]
    Version string // [--version=<version>]
    DryRun bool // [--dry-run]
    Insecure bool // [--insecure]
}

func (u Update) Args() []string {
    var args = []string{"plugin", "update"}
    args = utils.MakeArg(args, "[<plugin>...]", u.Plugin)
    args = utils.MakeArg(args, "[--all]", u.All)
    args = utils.MakeArg(args, "[--exclude=<name>]", u.Exclude)
    args = utils.MakeArg(args, "[--minor]", u.Minor)
    args = utils.MakeArg(args, "[--patch]", u.Patch)
    args = utils.MakeArg(args, "[--format=<format>]", u.Format)
    args = utils.MakeArg(args, "[--version=<version>]", u.Version)
    args = utils.MakeArg(args, "[--dry-run]", u.DryRun)
    args = utils.MakeArg(args, "[--insecure]", u.Insecure)
    return args
}

