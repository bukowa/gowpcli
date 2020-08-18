/*
## OPTIONS
	<db-id>
	: Database ID for the menu item.
	[--title=<title>]
	: Set a custom title for the menu item.
	[--link=<link>]
	: Set a custom url for the menu item.
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
## EXAMPLES
	    $ wp menu item update 45 --title=WordPress --link='http://wordpress.org' --target=_blank --position=2
	    Success: Menu item updated.
	
 */
package update
import utils "github.com/bukowa/gowpcli"

// Update //Updates a menu item.
type Update struct {
    DbId string // <db-id>
    Title string // [--title=<title>]
    Link string // [--link=<link>]
    Description string // [--description=<description>]
    AttrTitle string // [--attr-title=<attr-title>]
    Target string // [--target=<target>]
    Classes string // [--classes=<classes>]
    Position string // [--position=<position>]
    ParentId string // [--parent-id=<parent-id>]
}

func (u Update) Args() []string {
    var args = []string{"menu", "item", "update"}
    args = utils.MakeArg(args, "<db-id>", u.DbId)
    args = utils.MakeArg(args, "[--title=<title>]", u.Title)
    args = utils.MakeArg(args, "[--link=<link>]", u.Link)
    args = utils.MakeArg(args, "[--description=<description>]", u.Description)
    args = utils.MakeArg(args, "[--attr-title=<attr-title>]", u.AttrTitle)
    args = utils.MakeArg(args, "[--target=<target>]", u.Target)
    args = utils.MakeArg(args, "[--classes=<classes>]", u.Classes)
    args = utils.MakeArg(args, "[--position=<position>]", u.Position)
    args = utils.MakeArg(args, "[--parent-id=<parent-id>]", u.ParentId)
    return args
}

