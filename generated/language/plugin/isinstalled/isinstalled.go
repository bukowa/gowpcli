/*
## INFO
	Returns exit code 0 when installed, 1 when uninstalled.
## OPTIONS
	<plugin>
	: Plugin to check for.
	<language>...
	: The language code to check.
## EXAMPLES
	    # Check whether the German language is installed for Akismet; exit status 0 if installed, otherwise 1.
	    $ wp language plugin is-installed akismet de_DE
	    $ echo $?
	    1
	
 */
package isinstalled
import utils "github.com/bukowa/gowpcli"

// IsInstalled //Checks if a given language is installed.
type IsInstalled struct {
    Plugin string // <plugin>
    Language []string // <language>...
}

func (i IsInstalled) Args() []string {
    var args = []string{"language", "plugin", "is-installed"}
    args = utils.MakeArg(args, "<plugin>", i.Plugin)
    args = utils.MakeArg(args, "<language>...", i.Language)
    return args
}

