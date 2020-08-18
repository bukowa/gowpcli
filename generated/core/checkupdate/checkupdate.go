/*
## INFO
	Lists the most recent versions when there are updates available,
	or success message when up to date.
## OPTIONS
	[--minor]
	: Compare only the first two parts of the version number.
	[--major]
	: Compare only the first part of the version number.
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
	  - count
	  - json
	  - yaml
	---
## EXAMPLES
	    $ wp core check-update
	    +---------+-------------+-------------------------------------------------------------+
	    | version | update_type | package_url                                                 |
	    +---------+-------------+-------------------------------------------------------------+
	    | 4.5.2   | major       | https://downloads.wordpress.org/release/wordpress-4.5.2.zip |
	    +---------+-------------+-------------------------------------------------------------+
	
 */
package checkupdate
import utils "github.com/bukowa/gowpcli"

// CheckUpdate //Checks for WordPress updates via Version Check API.
type CheckUpdate struct {
    Minor bool // [--minor]
    Major bool // [--major]
    Field string // [--field=<field>]
    Fields string // [--fields=<fields>]
    Format string // [--format=<format>]
}

func (c CheckUpdate) Args() []string {
    var args = []string{"core", "check-update"}
    args = utils.MakeArg(args, "[--minor]", c.Minor)
    args = utils.MakeArg(args, "[--major]", c.Major)
    args = utils.MakeArg(args, "[--field=<field>]", c.Field)
    args = utils.MakeArg(args, "[--fields=<fields>]", c.Fields)
    args = utils.MakeArg(args, "[--format=<format>]", c.Format)
    return args
}

