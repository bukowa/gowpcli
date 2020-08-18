/*
## INFO
	Displays a list of the plugins installed on the site with activation
	status, whether or not there's an update available, etc.
	Use `--status=dropin` to list installed dropins (e.g. `object-cache.php`).
## OPTIONS
	[--<field>=<value>]
	: Filter results based on the value of a field.
	[--field=<field>]
	: Prints the value of a single field for each plugin.
	[--fields=<fields>]
	: Limit the output to specific object fields.
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
## AVAILABLE FIELDS
	These fields will be displayed by default for each plugin:
	* name
	* status
	* update
	* version
	These fields are optionally available:
	* update_version
	* update_package
	* update_id
	* title
	* description
	* file
## EXAMPLES
	    # List active plugins on the site.
	    $ wp plugin list --status=active --format=json
	    [{"name":"dynamic-hostname","status":"active","update":"none","version":"0.4.2"},{"name":"tinymce-templates","status":"active","update":"none","version":"4.4.3"},{"name":"wp-multibyte-patch","status":"active","update":"none","version":"2.4"},{"name":"wp-total-hacks","status":"active","update":"none","version":"2.0.1"}]
	    # List plugins on each site in a network.
	    $ wp site list --field=url | xargs -I % wp plugin list --url=%
	    +---------+----------------+--------+---------+
	    | name    | status         | update | version |
	    +---------+----------------+--------+---------+
	    | akismet | active-network | none   | 3.1.11  |
	    | hello   | inactive       | none   | 1.6     |
	    +---------+----------------+--------+---------+
	    +---------+----------------+--------+---------+
	    | name    | status         | update | version |
	    +---------+----------------+--------+---------+
	    | akismet | active-network | none   | 3.1.11  |
	    | hello   | inactive       | none   | 1.6     |
	    +---------+----------------+--------+---------+
	
 */
package list
import utils "github.com/bukowa/gowpcli"

// List //Gets a list of plugins.
type List struct {
    FieldMap map[string]string // [--<field>=<value>]
    Field string // [--field=<field>]
    Fields string // [--fields=<fields>]
    Format string // [--format=<format>]
}

func (l List) Args() []string {
    var args = []string{"plugin", "list"}
    args = utils.MakeArg(args, "[--<field>=<value>]", l.FieldMap)
    args = utils.MakeArg(args, "[--field=<field>]", l.Field)
    args = utils.MakeArg(args, "[--fields=<fields>]", l.Fields)
    args = utils.MakeArg(args, "[--format=<format>]", l.Format)
    return args
}

