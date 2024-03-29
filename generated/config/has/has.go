/*
## OPTIONS
	<name>
	: Name of the wp-config.php constant or variable.
	[--type=<type>]
	: Type of the config value to set. Defaults to 'all'.
	---
	default: all
	options:
	  - constant
	  - variable
	  - all
	---
	[--config-file=<path>]
	: Specify the file path to the config file to be checked. Defaults to the root of the
	WordPress installation and the filename "wp-config.php".
## EXAMPLES
	    # Check whether the DB_PASSWORD constant exists in the wp-config.php file.
	    $ wp config has DB_PASSWORD
	
 */
package has
import utils "github.com/bukowa/gowpcli"

// Has //Checks whether a specific constant or variable exists in the wp-config.php file.
type Has struct {
    Name string // <name>
    Type string // [--type=<type>]
    ConfigFile string // [--config-file=<path>]
}

func (h Has) Args() []string {
    var args = []string{"config", "has"}
    args = utils.MakeArg(args, "<name>", h.Name)
    args = utils.MakeArg(args, "[--type=<type>]", h.Type)
    args = utils.MakeArg(args, "[--config-file=<path>]", h.ConfigFile)
    return args
}

