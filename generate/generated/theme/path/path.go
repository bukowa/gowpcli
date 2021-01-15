/*
## OPTIONS
	[<theme>]
	: The theme to get the path to. Path includes "style.css" file.
	If not set, will return the path to the themes directory.
	[--dir]
	: If set, get the path to the closest parent directory, instead of the
	theme's "style.css" file.
## EXAMPLES
	    # Get theme path
	    $ wp theme path
	    /var/www/example.com/public_html/wp-content/themes
	    # Change directory to theme path
	    $ cd $(wp theme path)
	
 */
package path
import utils "github.com/bukowa/gowpcli"

// Path //Gets the path to a theme or to the theme directory.
type Path struct {
    Theme string // [<theme>]
    Dir bool // [--dir]
}

func (p Path) Args() []string {
    var args = []string{"theme", "path"}
    args = utils.MakeArg(args, "[<theme>]", p.Theme)
    args = utils.MakeArg(args, "[--dir]", p.Dir)
    return args
}

