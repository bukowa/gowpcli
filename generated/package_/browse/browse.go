/*
## INFO
	Lists packages available for installation from the [Package Index](http://wp-cli.org/package-index/).
	Although the package index will remain in place for backward compatibility reasons, it has been
	deprecated and will not be updated further. Please refer to https://github.com/wp-cli/ideas/issues/51
	to read about its potential replacement.
## OPTIONS
	[--fields=<fields>]
	: Limit the output to specific fields. Defaults to all fields.
	[--format=<format>]
	: Render output in a particular format.
	---
	default: table
	options:
	  - table
	  - csv
	  - ids
	  - json
	  - yaml
	---
## AVAILABLE FIELDS
	These fields will be displayed by default for each package:
	* name
	* description
	* authors
	* version
	There are no optionally available fields.
## EXAMPLES
	    $ wp package browse --format=yaml
	    ---
	    10up/mu-migration:
	      name: 10up/mu-migration
	      description: A set of WP-CLI commands to support the migration of single WordPress instances to multisite
	      authors: Nícholas André
	      version: dev-master, dev-develop
	    aaemnnosttv/wp-cli-dotenv-command:
	      name: aaemnnosttv/wp-cli-dotenv-command
	      description: Dotenv commands for WP-CLI
	      authors: Evan Mattson
	      version: v0.1, v0.1-beta.1, v0.2, dev-master, dev-dev, dev-develop, dev-tests/behat
	    aaemnnosttv/wp-cli-http-command:
	      name: aaemnnosttv/wp-cli-http-command
	      description: WP-CLI command for using the WordPress HTTP API
	      authors: Evan Mattson
	      version: dev-master
	
 */
package browse
import utils "github.com/bukowa/gowpcli"

// Browse //Browses WP-CLI packages available for installation.
type Browse struct {
    Fields string // [--fields=<fields>]
    Format string // [--format=<format>]
}

func (b Browse) Args() []string {
    var args = []string{"package", "browse"}
    args = utils.MakeArg(args, "[--fields=<fields>]", b.Fields)
    args = utils.MakeArg(args, "[--format=<format>]", b.Format)
    return args
}

