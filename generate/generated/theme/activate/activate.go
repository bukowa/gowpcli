/*
## OPTIONS
	<theme>
	: The theme to activate.
## EXAMPLES
	    $ wp theme activate twentysixteen
	    Success: Switched to 'Twenty Sixteen' theme.
	
 */
package activate
import utils "github.com/bukowa/gowpcli"

// Activate //Activates a theme.
type Activate struct {
    Theme string // <theme>
}

func (a Activate) Args() []string {
    var args = []string{"theme", "activate"}
    args = utils.MakeArg(args, "<theme>", a.Theme)
    return args
}

