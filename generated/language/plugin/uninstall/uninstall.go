/*
## OPTIONS
	<plugin>
	: Plugin to uninstall language for.
	<language>...
	: Language code to uninstall.
## EXAMPLES
	    $ wp language plugin uninstall hello-dolly ja
	    Success: Language uninstalled.
	
 */
package uninstall
import utils "github.com/bukowa/gowpcli"

// Uninstall //Uninstalls a given language for a plugin.
type Uninstall struct {
    Plugin string // <plugin>
    Language []string // <language>...
}

func (u Uninstall) Args() []string {
    var args = []string{"language", "plugin", "uninstall"}
    args = utils.MakeArg(args, "<plugin>", u.Plugin)
    args = utils.MakeArg(args, "<language>...", u.Language)
    return args
}

