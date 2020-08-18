/*
## OPTIONS
	[<plugin>]
	: The plugin to get the path to. If not set, will return the path to the
	plugins directory.
	[--dir]
	: If set, get the path to the closest parent directory, instead of the
	plugin file.
## EXAMPLES
	    $ cd $(wp plugin path) && pwd
	    /var/www/wordpress/wp-content/plugins
	
 */
package path
import utils "github.com/bukowa/gowpcli"

// Path //Gets the path to a plugin or to the plugin directory.
type Path struct {
    Plugin string // [<plugin>]
    Dir bool // [--dir]
}

func (p Path) Args() []string {
    var args = []string{"plugin", "path"}
    args = utils.MakeArg(args, "[<plugin>]", p.Plugin)
    args = utils.MakeArg(args, "[--dir]", p.Dir)
    return args
}

