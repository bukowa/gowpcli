/*
## OPTIONS
	<theme>
	: Theme to uninstall language for.
	<language>...
	: Language code to uninstall.
## EXAMPLES
	    $ wp language theme uninstall twentyten ja
	    Success: Language uninstalled.
	
 */
package uninstall
import utils "github.com/bukowa/gowpcli"

// Uninstall //Uninstalls a given language for a theme.
type Uninstall struct {
    Theme string // <theme>
    Language []string // <language>...
}

func (u Uninstall) Args() []string {
    var args = []string{"language", "theme", "uninstall"}
    args = utils.MakeArg(args, "<theme>", u.Theme)
    args = utils.MakeArg(args, "<language>...", u.Language)
    return args
}

