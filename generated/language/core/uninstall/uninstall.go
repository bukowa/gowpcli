/*
## OPTIONS
	<language>...
	: Language code to uninstall.
## EXAMPLES
	    $ wp language core uninstall ja
	    Success: Language uninstalled.
	
 */
package uninstall
import utils "github.com/bukowa/gowpcli"

// Uninstall //Uninstalls a given language.
type Uninstall struct {
    Language []string // <language>...
}

func (u Uninstall) Args() []string {
    var args = []string{"language", "core", "uninstall"}
    args = utils.MakeArg(args, "<language>...", u.Language)
    return args
}

