/*
## OPTIONS
	[<mod>...]
	: One or more mods to remove.
	[--all]
	: Remove all theme mods.
## EXAMPLES
	    # Remove all theme mods.
	    $ wp theme mod remove --all
	    Success: Theme mods removed.
	    # Remove single theme mod.
	    $ wp theme mod remove background_color
	    Success: 1 mod removed.
	    # Remove multiple theme mods.
	    $ wp theme mod remove background_color header_textcolor
	    Success: 2 mods removed.
	
 */
package remove
import utils "github.com/bukowa/gowpcli"

// Remove //Removes one or more theme mods.
type Remove struct {
    Mod []string // [<mod>...]
    All bool // [--all]
}

func (r Remove) Args() []string {
    var args = []string{"theme", "mod", "remove"}
    args = utils.MakeArg(args, "[<mod>...]", r.Mod)
    args = utils.MakeArg(args, "[--all]", r.All)
    return args
}

