/*
## INFO
	**Warning: `wp scaffold block` is deprecated.**
	The official script to generate a block is the [@wordpress/create-block](https://developer.wordpress.org/block-editor/designers-developers/developers/packages/packages-create-block/) package.
	See the [Create a Block tutorial](https://developer.wordpress.org/block-editor/getting-started/create-block/) for a complete walk-through.
## OPTIONS
	<slug>
	: The internal name of the block.
	[--title=<title>]
	: The display title for your block.
	[--dashicon=<dashicon>]
	: The dashicon to make it easier to identify your block.
	[--category=<category>]
	: The category name to help users browse and discover your block.
	---
	default: widgets
	options:
	  - common
	  - embed
	  - formatting
	  - layout
	  - widgets
	---
	[--theme]
	: Create files in the active theme directory. Specify a theme with `--theme=<theme>` to have the file placed in that theme.
	[--plugin=<plugin>]
	: Create files in the given plugin's directory.
	[--force]
	: Overwrite files that already exist.
	
 */
package block
import utils "github.com/bukowa/gowpcli"

// Block //Generates PHP, JS and CSS code for registering a Gutenberg block for a plugin or theme.
type Block struct {
    Slug string // <slug>
    Title string // [--title=<title>]
    Dashicon string // [--dashicon=<dashicon>]
    Category string // [--category=<category>]
    Theme bool // [--theme]
    Plugin string // [--plugin=<plugin>]
    Force bool // [--force]
}

func (b Block) Args() []string {
    var args = []string{"scaffold", "block"}
    args = utils.MakeArg(args, "<slug>", b.Slug)
    args = utils.MakeArg(args, "[--title=<title>]", b.Title)
    args = utils.MakeArg(args, "[--dashicon=<dashicon>]", b.Dashicon)
    args = utils.MakeArg(args, "[--category=<category>]", b.Category)
    args = utils.MakeArg(args, "[--theme]", b.Theme)
    args = utils.MakeArg(args, "[--plugin=<plugin>]", b.Plugin)
    args = utils.MakeArg(args, "[--force]", b.Force)
    return args
}

