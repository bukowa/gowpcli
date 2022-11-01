/*
## OPTIONS
	<name>
	: Name of the wp-config.php constant or variable.
	[--type=<type>]
	: Type of the config value to delete. Defaults to 'all'.
	---
	default: all
	options:
	  - constant
	  - variable
	  - all
	---
	[--config-file=<path>]
	: Specify the file path to the config file to be modified. Defaults to the root of the
	WordPress installation and the filename "wp-config.php".
## EXAMPLES
	    # Delete the COOKIE_DOMAIN constant from the wp-config.php file.
	    $ wp config delete COOKIE_DOMAIN
	
 */
package delete
import utils "github.com/bukowa/gowpcli"

// Delete //Deletes a specific constant or variable from the wp-config.php file.
type Delete struct {
    Name string // <name>
    Type string // [--type=<type>]
    ConfigFile string // [--config-file=<path>]
}

func (d Delete) Args() []string {
    var args = []string{"config", "delete"}
    args = utils.MakeArg(args, "<name>", d.Name)
    args = utils.MakeArg(args, "[--type=<type>]", d.Type)
    args = utils.MakeArg(args, "[--config-file=<path>]", d.ConfigFile)
    return args
}

