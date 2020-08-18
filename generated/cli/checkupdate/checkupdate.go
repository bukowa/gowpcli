/*
## INFO
	Queries the Github releases API. Returns available versions if there are
	updates available, or success message if using the latest release.
## OPTIONS
	[--patch]
	: Only list patch updates.
	[--minor]
	: Only list minor updates.
	[--major]
	: Only list major updates.
	[--field=<field>]
	: Prints the value of a single field for each update.
	[--fields=<fields>]
	: Limit the output to specific object fields. Defaults to version,update_type,package_url.
	[--format=<format>]
	: Render output in a particular format.
	---
	default: table
	options:
	  - table
	  - csv
	  - json
	  - count
	  - yaml
	---
## EXAMPLES
	    # Check for update.
	    $ wp cli check-update
	    Success: WP-CLI is at the latest version.
	    # Check for update and new version is available.
	    $ wp cli check-update
	    +---------+-------------+-------------------------------------------------------------------------------+
	    | version | update_type | package_url                                                                   |
	    +---------+-------------+-------------------------------------------------------------------------------+
	    | 0.24.1  | patch       | https://github.com/wp-cli/wp-cli/releases/download/v0.24.1/wp-cli-0.24.1.phar |
	    +---------+-------------+-------------------------------------------------------------------------------+
	
 */
package checkupdate
import utils "github.com/bukowa/gowpcli"

// CheckUpdate //Checks to see if there is a newer version of WP-CLI available.
type CheckUpdate struct {
    Patch bool // [--patch]
    Minor bool // [--minor]
    Major bool // [--major]
    Field string // [--field=<field>]
    Fields string // [--fields=<fields>]
    Format string // [--format=<format>]
}

func (c CheckUpdate) Args() []string {
    var args = []string{"cli", "check-update"}
    args = utils.MakeArg(args, "[--patch]", c.Patch)
    args = utils.MakeArg(args, "[--minor]", c.Minor)
    args = utils.MakeArg(args, "[--major]", c.Major)
    args = utils.MakeArg(args, "[--field=<field>]", c.Field)
    args = utils.MakeArg(args, "[--fields=<fields>]", c.Fields)
    args = utils.MakeArg(args, "[--format=<format>]", c.Format)
    return args
}

