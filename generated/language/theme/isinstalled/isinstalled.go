/*
## INFO
	Returns exit code 0 when installed, 1 when uninstalled.
## OPTIONS
	<theme>
	: Theme to check for.
	<language>...
	: The language code to check.
## EXAMPLES
	    # Check whether the German language is installed for Twenty Seventeen; exit status 0 if installed, otherwise 1.
	    $ wp language theme is-installed twentyseventeen de_DE
	    $ echo $?
	    1
	
 */
package isinstalled
import utils "github.com/bukowa/gowpcli"

// IsInstalled //Checks if a given language is installed.
type IsInstalled struct {
    Theme string // <theme>
    Language []string // <language>...
}

func (i IsInstalled) Args() []string {
    var args = []string{"language", "theme", "is-installed"}
    args = utils.MakeArg(args, "<theme>", i.Theme)
    args = utils.MakeArg(args, "<language>...", i.Language)
    return args
}

