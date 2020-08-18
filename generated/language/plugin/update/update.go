/*
## OPTIONS
	[<plugin>...]
	: One or more plugins to update languages for.
	[--all]
	: If set, languages for all plugins will be updated.
	[--dry-run]
	: Preview which translations would be updated.
## EXAMPLES
	    $ wp language plugin update --all
	    Updating 'Japanese' translation for Akismet 3.1.11...
	    Downloading translation from https://downloads.wordpress.org/translation/plugin/akismet/3.1.11/ja.zip...
	    Translation updated successfully.
	    Success: Updated 1/1 translation.
	
 */
package update
import utils "github.com/bukowa/gowpcli"

// Update //Updates installed languages for one or more plugins.
type Update struct {
    Plugin []string // [<plugin>...]
    All bool // [--all]
    DryRun bool // [--dry-run]
}

func (u Update) Args() []string {
    var args = []string{"language", "plugin", "update"}
    args = utils.MakeArg(args, "[<plugin>...]", u.Plugin)
    args = utils.MakeArg(args, "[--all]", u.All)
    args = utils.MakeArg(args, "[--dry-run]", u.DryRun)
    return args
}

