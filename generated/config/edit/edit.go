/*
## OPTIONS
	[--config-file=<path>]
	: Specify the file path to the config file to be edited. Defaults to the root of the
	WordPress installation and the filename "wp-config.php".
## EXAMPLES
	    # Launch system editor to edit wp-config.php file
	    $ wp config edit
	    # Edit wp-config.php file in a specific editor
	    $ EDITOR=vim wp config edit
	
 */
package edit
import utils "github.com/bukowa/gowpcli"

// Edit //Launches system editor to edit the wp-config.php file.
type Edit struct {
    ConfigFile string // [--config-file=<path>]
}

func (e Edit) Args() []string {
    var args = []string{"config", "edit"}
    args = utils.MakeArg(args, "[--config-file=<path>]", e.ConfigFile)
    return args
}

