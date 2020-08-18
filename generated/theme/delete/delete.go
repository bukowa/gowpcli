/*
## INFO
	Removes the theme or themes from the filesystem.
## OPTIONS
	[<theme>...]
	: One or more themes to delete.
	[--all]
	: If set, all themes will be deleted except active theme.
	[--force]
	: To delete active theme use this.
## EXAMPLES
	    $ wp theme delete twentytwelve
	    Deleted 'twentytwelve' theme.
	    Success: Deleted 1 of 1 themes.
	
 */
package delete
import utils "github.com/bukowa/gowpcli"

// Delete //Deletes one or more themes.
type Delete struct {
    Theme []string // [<theme>...]
    All bool // [--all]
    Force bool // [--force]
}

func (d Delete) Args() []string {
    var args = []string{"theme", "delete"}
    args = utils.MakeArg(args, "[<theme>...]", d.Theme)
    args = utils.MakeArg(args, "[--all]", d.All)
    args = utils.MakeArg(args, "[--force]", d.Force)
    return args
}

