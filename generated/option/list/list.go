/*
## OPTIONS
	[--search=<pattern>]
	: Use wildcards ( * and ? ) to match option name.
	[--exclude=<pattern>]
	: Pattern to exclude. Use wildcards ( * and ? ) to match option name.
	[--autoload=<value>]
	: Match only autoload options when value is on, and only not-autoload option when off.
	[--transients]
	: List only transients. Use `--no-transients` to ignore all transients.
	[--unserialize]
	: Unserialize option values in output.
	[--field=<field>]
	: Prints the value of a single field.
	[--fields=<fields>]
	: Limit the output to specific object fields.
	[--format=<format>]
	: The serialization format for the value. total_bytes displays the total size of matching options in bytes.
	---
	default: table
	options:
	  - table
	  - json
	  - csv
	  - count
	  - yaml
	  - total_bytes
	---
	[--orderby=<fields>]
	: Set orderby which field.
	---
	default: option_id
	options:
	 - option_id
	 - option_name
	 - option_value
	---
	[--order=<order>]
	: Set ascending or descending order.
	---
	default: asc
	options:
	 - asc
	 - desc
	---
## AVAILABLE FIELDS
	This field will be displayed by default for each matching option:
	* option_name
	* option_value
	These fields are optionally available:
	* autoload
	* size_bytes
## EXAMPLES
	    # Get the total size of all autoload options.
	    $ wp option list --autoload=on --format=total_bytes
	    33198
	    # Find biggest transients.
	    $ wp option list --search="*_transient_*" --fields=option_name,size_bytes | sort -n -k 2 | tail
	    option_name size_bytes
	    _site_transient_timeout_theme_roots 10
	    _site_transient_theme_roots 76
	    _site_transient_update_themes   181
	    _site_transient_update_core 808
	    _site_transient_update_plugins  6645
	    # List all options beginning with "i2f_".
	    $ wp option list --search="i2f_*"
	    +-------------+--------------+
	    | option_name | option_value |
	    +-------------+--------------+
	    | i2f_version | 0.1.0        |
	    +-------------+--------------+
	    # Delete all options beginning with "theme_mods_".
	    $ wp option list --search="theme_mods_*" --field=option_name | xargs -I % wp option delete %
	    Success: Deleted 'theme_mods_twentysixteen' option.
	    Success: Deleted 'theme_mods_twentyfifteen' option.
	    Success: Deleted 'theme_mods_twentyfourteen' option.
	
 */
package list
import utils "github.com/bukowa/gowpcli"

// List //Lists options and their values.
type List struct {
    Search string // [--search=<pattern>]
    Exclude string // [--exclude=<pattern>]
    Autoload string // [--autoload=<value>]
    Transients bool // [--transients]
    Unserialize bool // [--unserialize]
    Field string // [--field=<field>]
    Fields string // [--fields=<fields>]
    Format string // [--format=<format>]
    Orderby string // [--orderby=<fields>]
    Order string // [--order=<order>]
}

func (l List) Args() []string {
    var args = []string{"option", "list"}
    args = utils.MakeArg(args, "[--search=<pattern>]", l.Search)
    args = utils.MakeArg(args, "[--exclude=<pattern>]", l.Exclude)
    args = utils.MakeArg(args, "[--autoload=<value>]", l.Autoload)
    args = utils.MakeArg(args, "[--transients]", l.Transients)
    args = utils.MakeArg(args, "[--unserialize]", l.Unserialize)
    args = utils.MakeArg(args, "[--field=<field>]", l.Field)
    args = utils.MakeArg(args, "[--fields=<fields>]", l.Fields)
    args = utils.MakeArg(args, "[--format=<format>]", l.Format)
    args = utils.MakeArg(args, "[--orderby=<fields>]", l.Orderby)
    args = utils.MakeArg(args, "[--order=<order>]", l.Order)
    return args
}

