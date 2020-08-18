package addpost

//Adds a post as a menu item.

type AddPost struct {
    
    Menu string // <menu>
    
    PostId string // <post-id>
    
    Title string // [--title=<title>]
    
    Link string // [--link=<link>]
    
    Description string // [--description=<description>]
    
    AttrTitle string // [--attr-title=<attr-title>]
    
    Target string // [--target=<target>]
    
    Classes string // [--classes=<classes>]
    
    Position string // [--position=<position>]
    
    ParentId string // [--parent-id=<parent-id>]
    
    Porcelain bool // [--porcelain]
    
}

//## OPTIONS
//
//<menu>
//: The name, slug, or term ID for the menu.
//
//<post-id>
//: Post ID to add to the menu.
//
//[--title=<title>]
//: Set a custom title for the menu item.
//
//[--link=<link>]
//: Set a custom url for the menu item.
//
//[--description=<description>]
//: Set a custom description for the menu item.
//
//[--attr-title=<attr-title>]
//: Set a custom title attribute for the menu item.
//
//[--target=<target>]
//: Set a custom link target for the menu item.
//
//[--classes=<classes>]
//: Set a custom link classes for the menu item.
//
//[--position=<position>]
//: Specify the position of this menu item.
//
//[--parent-id=<parent-id>]
//: Make this menu item a child of another menu item.
//
//[--porcelain]
//: Output just the new menu item id.
//
//## EXAMPLES
//
//    $ wp menu item add-post sidebar-menu 33 --title="Custom Test Post"
//    Success: Menu item added.
//
//