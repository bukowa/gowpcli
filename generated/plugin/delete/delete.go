/*
## OPTIONS
	[<plugin>...]
	: One or more plugins to delete.
	[--all]
	: If set, all plugins will be deleted.
## EXAMPLES
	    # Delete plugin
	    $ wp plugin delete hello
	    Deleted 'hello' plugin.
	    Success: Deleted 1 of 1 plugins.
	    # Delete inactive plugins
	    $ wp plugin delete $(wp plugin list --status=inactive --field=name)
	    Deleted 'tinymce-templates' plugin.
	    Success: Deleted 1 of 1 plugins.
	
 */
package delete
import utils "github.com/bukowa/gowpcli"

// Delete //Deletes plugin files without deactivating or uninstalling.
type Delete struct {
    Plugin []string // [<plugin>...]
    All bool // [--all]
}

func (d Delete) Args() []string {
    var args = []string{"plugin", "delete"}
    args = utils.MakeArg(args, "[<plugin>...]", d.Plugin)
    args = utils.MakeArg(args, "[--all]", d.All)
    return args
}

