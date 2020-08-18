/*
## OPTIONS
	<slug>
	: The internal name of the post type.
	[--label=<label>]
	: The text used to translate the update messages.
	[--textdomain=<textdomain>]
	: The textdomain to use for the labels.
	[--dashicon=<dashicon>]
	: The dashicon to use in the menu.
	[--theme]
	: Create a file in the active theme directory, instead of sending to
	STDOUT. Specify a theme with `--theme=<theme>` to have the file placed in that theme.
	[--plugin=<plugin>]
	: Create a file in the given plugin's directory, instead of sending to STDOUT.
	[--raw]
	: Just generate the `register_post_type()` call and nothing else.
	[--force]
	: Overwrite files that already exist.
## EXAMPLES
	    # Generate a 'movie' post type for the 'simple-life' theme
	    $ wp scaffold post-type movie --label=Movie --theme=simple-life
	    Success: Created '/var/www/example.com/public_html/wp-content/themes/simple-life/post-types/movie.php'.
	
 */
package posttype
import utils "github.com/bukowa/gowpcli"

// PostType //Generates PHP code for registering a custom post type.
type PostType struct {
    Slug string // <slug>
    Label string // [--label=<label>]
    Textdomain string // [--textdomain=<textdomain>]
    Dashicon string // [--dashicon=<dashicon>]
    Theme bool // [--theme]
    Plugin string // [--plugin=<plugin>]
    Raw bool // [--raw]
    Force bool // [--force]
}

func (p PostType) Args() []string {
    var args = []string{"scaffold", "post-type"}
    args = utils.MakeArg(args, "<slug>", p.Slug)
    args = utils.MakeArg(args, "[--label=<label>]", p.Label)
    args = utils.MakeArg(args, "[--textdomain=<textdomain>]", p.Textdomain)
    args = utils.MakeArg(args, "[--dashicon=<dashicon>]", p.Dashicon)
    args = utils.MakeArg(args, "[--theme]", p.Theme)
    args = utils.MakeArg(args, "[--plugin=<plugin>]", p.Plugin)
    args = utils.MakeArg(args, "[--raw]", p.Raw)
    args = utils.MakeArg(args, "[--force]", p.Force)
    return args
}

