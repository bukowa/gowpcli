/*
## INFO
	See the [Underscores website](https://underscores.me/) for more details.
## OPTIONS
	<slug>
	: The slug for the new theme, used for prefixing functions.
	[--activate]
	: Activate the newly downloaded theme.
	[--enable-network]
	: Enable the newly downloaded theme for the entire network.
	[--theme_name=<title>]
	: What to put in the 'Theme Name:' header in 'style.css'.
	[--author=<full-name>]
	: What to put in the 'Author:' header in 'style.css'.
	[--author_uri=<uri>]
	: What to put in the 'Author URI:' header in 'style.css'.
	[--sassify]
	: Include stylesheets as SASS.
	[--woocommerce]
	: Include WooCommerce boilerplate files.
	[--force]
	: Overwrite files that already exist.
## EXAMPLES
	    # Generate a theme with name "Sample Theme" and author "John Doe"
	    $ wp scaffold _s sample-theme --theme_name="Sample Theme" --author="John Doe"
	    Success: Created theme 'Sample Theme'.
	
 */
package underscores
import utils "github.com/bukowa/gowpcli"

// Underscores //Generates starter code for a theme based on _s.
type Underscores struct {
    Slug string // <slug>
    Activate bool // [--activate]
    EnableNetwork bool // [--enable-network]
    ThemeName string // [--theme_name=<title>]
    Author string // [--author=<full-name>]
    AuthorUri string // [--author_uri=<uri>]
    Sassify bool // [--sassify]
    Woocommerce bool // [--woocommerce]
    Force bool // [--force]
}

func (u Underscores) Args() []string {
    var args = []string{"scaffold", "underscores"}
    args = utils.MakeArg(args, "<slug>", u.Slug)
    args = utils.MakeArg(args, "[--activate]", u.Activate)
    args = utils.MakeArg(args, "[--enable-network]", u.EnableNetwork)
    args = utils.MakeArg(args, "[--theme_name=<title>]", u.ThemeName)
    args = utils.MakeArg(args, "[--author=<full-name>]", u.Author)
    args = utils.MakeArg(args, "[--author_uri=<uri>]", u.AuthorUri)
    args = utils.MakeArg(args, "[--sassify]", u.Sassify)
    args = utils.MakeArg(args, "[--woocommerce]", u.Woocommerce)
    args = utils.MakeArg(args, "[--force]", u.Force)
    return args
}

