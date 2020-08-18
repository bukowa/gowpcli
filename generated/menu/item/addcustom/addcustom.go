/*
## OPTIONS
	<menu>
	: The name, slug, or term ID for the menu.
	<title>
	: Title for the link.
	<link>
	: Target URL for the link.
	[--description=<description>]
	: Set a custom description for the menu item.
	[--attr-title=<attr-title>]
	: Set a custom title attribute for the menu item.
	[--target=<target>]
	: Set a custom link target for the menu item.
	[--classes=<classes>]
	: Set a custom link classes for the menu item.
	[--position=<position>]
	: Specify the position of this menu item.
	[--parent-id=<parent-id>]
	: Make this menu item a child of another menu item.
	[--porcelain]
	: Output just the new menu item id.
## EXAMPLES
	    $ wp menu item add-custom sidebar-menu Apple http://apple.com
	    Success: Menu item added.
	
 */
package addcustom
import utils "github.com/bukowa/gowpcli"

// AddCustom //Adds a custom menu item.
type AddCustom struct {
    Menu string // <menu>
    Title string // <title>
    Link string // <link>
    Description string // [--description=<description>]
    AttrTitle string // [--attr-title=<attr-title>]
    Target string // [--target=<target>]
    Classes string // [--classes=<classes>]
    Position string // [--position=<position>]
    ParentId string // [--parent-id=<parent-id>]
    Porcelain bool // [--porcelain]
}

func (a AddCustom) Args() []string {
    var args = []string{"menu", "item", "add-custom"}
    args = utils.MakeArg(args, "<menu>", a.Menu)
    args = utils.MakeArg(args, "<title>", a.Title)
    args = utils.MakeArg(args, "<link>", a.Link)
    args = utils.MakeArg(args, "[--description=<description>]", a.Description)
    args = utils.MakeArg(args, "[--attr-title=<attr-title>]", a.AttrTitle)
    args = utils.MakeArg(args, "[--target=<target>]", a.Target)
    args = utils.MakeArg(args, "[--classes=<classes>]", a.Classes)
    args = utils.MakeArg(args, "[--position=<position>]", a.Position)
    args = utils.MakeArg(args, "[--parent-id=<parent-id>]", a.ParentId)
    args = utils.MakeArg(args, "[--porcelain]", a.Porcelain)
    return args
}

