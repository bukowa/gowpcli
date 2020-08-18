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
}

func (d Delete) Args() []string {
    var args = []string{"config", "delete"}
    args = utils.MakeArg(args, "<name>", d.Name)
    args = utils.MakeArg(args, "[--type=<type>]", d.Type)
    return args
}

