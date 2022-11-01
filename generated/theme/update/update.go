/*
## OPTIONS
	[<theme>...]
	: One or more themes to update.
	[--all]
	: If set, all themes that have updates will be updated.
	[--exclude=<theme-names>]
	: Comma separated list of theme names that should be excluded from updating.
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
	: If set, the theme will be updated to the specified version.
	[--dry-run]
	: Preview which themes would be updated.
	[--insecure]
	: Retry downloads without certificate validation if TLS handshake fails. Note: This makes the request vulnerable to a MITM attack.
## EXAMPLES
	    # Update multiple themes
	    $ wp theme update twentyfifteen twentysixteen
	    Downloading update from https://downloads.wordpress.org/theme/twentyfifteen.1.5.zip...
	    Unpacking the update...
	    Installing the latest version...
	    Removing the old version of the theme...
	    Theme updated successfully.
	    Downloading update from https://downloads.wordpress.org/theme/twentysixteen.1.2.zip...
	    Unpacking the update...
	    Installing the latest version...
	    Removing the old version of the theme...
	    Theme updated successfully.
	    +---------------+-------------+-------------+---------+
	    | name          | old_version | new_version | status  |
	    +---------------+-------------+-------------+---------+
	    | twentyfifteen | 1.4         | 1.5         | Updated |
	    | twentysixteen | 1.1         | 1.2         | Updated |
	    +---------------+-------------+-------------+---------+
	    Success: Updated 2 of 2 themes.
	    # Exclude themes updates when bulk updating the themes
	    $ wp theme update --all --exclude=twentyfifteen
	    Downloading update from https://downloads.wordpress.org/theme/astra.1.0.5.1.zip...
	    Unpacking the update...
	    Installing the latest version...
	    Removing the old version of the theme...
	    Theme updated successfully.
	    Downloading update from https://downloads.wordpress.org/theme/twentyseventeen.1.2.zip...
	    Unpacking the update...
	    Installing the latest version...
	    Removing the old version of the theme...
	    Theme updated successfully.
	    +-----------------+----------+---------+----------------+
	    | name            | status   | version | update_version |
	    +-----------------+----------+---------+----------------+
	    | astra           | inactive | 1.0.1   | 1.0.5.1        |
	    | twentyseventeen | inactive | 1.1     | 1.2            |
	    +-----------------+----------+---------+----------------+
	    Success: Updated 2 of 2 themes.
	    # Update all themes
	    $ wp theme update --all
	
 */
package update
import utils "github.com/bukowa/gowpcli"

// Update //Updates one or more themes.
type Update struct {
    Theme []string // [<theme>...]
    All bool // [--all]
    Exclude string // [--exclude=<theme-names>]
    Format string // [--format=<format>]
    Version string // [--version=<version>]
    DryRun bool // [--dry-run]
    Insecure bool // [--insecure]
}

func (u Update) Args() []string {
    var args = []string{"theme", "update"}
    args = utils.MakeArg(args, "[<theme>...]", u.Theme)
    args = utils.MakeArg(args, "[--all]", u.All)
    args = utils.MakeArg(args, "[--exclude=<theme-names>]", u.Exclude)
    args = utils.MakeArg(args, "[--format=<format>]", u.Format)
    args = utils.MakeArg(args, "[--version=<version>]", u.Version)
    args = utils.MakeArg(args, "[--dry-run]", u.DryRun)
    args = utils.MakeArg(args, "[--insecure]", u.Insecure)
    return args
}

