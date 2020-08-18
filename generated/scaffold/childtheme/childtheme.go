/*
## INFO
	Creates a child theme folder with `functions.php` and `style.css` files.
## OPTIONS
	<slug>
	: The slug for the new child theme.
	--parent_theme=<slug>
	: What to put in the 'Template:' header in 'style.css'.
	[--theme_name=<title>]
	: What to put in the 'Theme Name:' header in 'style.css'.
	[--author=<full-name>]
	: What to put in the 'Author:' header in 'style.css'.
	[--author_uri=<uri>]
	: What to put in the 'Author URI:' header in 'style.css'.
	[--theme_uri=<uri>]
	: What to put in the 'Theme URI:' header in 'style.css'.
	[--activate]
	: Activate the newly created child theme.
	[--enable-network]
	: Enable the newly created child theme for the entire network.
	[--force]
	: Overwrite files that already exist.
## EXAMPLES
	    # Generate a 'sample-theme' child theme based on TwentySixteen
	    $ wp scaffold child-theme sample-theme --parent_theme=twentysixteen
	    Success: Created '/var/www/example.com/public_html/wp-content/themes/sample-theme'.
	
 */
package childtheme
import utils "github.com/bukowa/gowpcli"

// ChildTheme //Generates child theme based on an existing theme.
type ChildTheme struct {
    Slug string // <slug>
    ParentTheme string // --parent_theme=<slug>
    ThemeName string // [--theme_name=<title>]
    Author string // [--author=<full-name>]
    AuthorUri string // [--author_uri=<uri>]
    ThemeUri string // [--theme_uri=<uri>]
    Activate bool // [--activate]
    EnableNetwork bool // [--enable-network]
    Force bool // [--force]
}

func (c ChildTheme) Args() []string {
    var args = []string{"scaffold", "child-theme"}
    args = utils.MakeArg(args, "<slug>", c.Slug)
    args = utils.MakeArg(args, "--parent_theme=<slug>", c.ParentTheme)
    args = utils.MakeArg(args, "[--theme_name=<title>]", c.ThemeName)
    args = utils.MakeArg(args, "[--author=<full-name>]", c.Author)
    args = utils.MakeArg(args, "[--author_uri=<uri>]", c.AuthorUri)
    args = utils.MakeArg(args, "[--theme_uri=<uri>]", c.ThemeUri)
    args = utils.MakeArg(args, "[--activate]", c.Activate)
    args = utils.MakeArg(args, "[--enable-network]", c.EnableNetwork)
    args = utils.MakeArg(args, "[--force]", c.Force)
    return args
}

