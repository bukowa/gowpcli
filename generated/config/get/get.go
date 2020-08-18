/*
## OPTIONS
	<name>
	: Name of the wp-config.php constant or variable.
	[--type=<type>]
	: Type of config value to retrieve. Defaults to 'all'.
	---
	default: all
	options:
	  - constant
	  - variable
	  - all
	---
	[--format=<format>]
	: Get value in a particular format.
	---
	default: var_export
	options:
	  - var_export
	  - json
	  - yaml
	---
## EXAMPLES
	    # Get the table_prefix as defined in wp-config.php file.
	    $ wp config get table_prefix
	    wp_
	
 */
package get
import utils "github.com/bukowa/gowpcli"

// Get //Gets the value of a specific constant or variable defined in wp-config.php file.
type Get struct {
    Name string // <name>
    Type string // [--type=<type>]
    Format string // [--format=<format>]
}

func (g Get) Args() []string {
    var args = []string{"config", "get"}
    args = utils.MakeArg(args, "<name>", g.Name)
    args = utils.MakeArg(args, "[--type=<type>]", g.Type)
    args = utils.MakeArg(args, "[--format=<format>]", g.Format)
    return args
}

