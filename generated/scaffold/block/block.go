package block

//Generates PHP, JS and CSS code for registering a Gutenberg block for a plugin or theme.

type Block struct {
    
    Slug string // <slug>
    
    Title string // [--title=<title>]
    
    Dashicon string // [--dashicon=<dashicon>]
    
    Category string // [--category=<category>]
    
    Theme bool // [--theme]
    
    Plugin string // [--plugin=<plugin>]
    
    Force bool // [--force]
    
}

//Blocks are the fundamental element of the Gutenberg editor. They are the primary way in which plugins and themes can register their own functionality and extend the capabilities of the editor.
//
//Visit the [Gutenberg handbook](https://wordpress.org/gutenberg/handbook/block-api/) to learn more about Block API.
//
//When you scaffold a block you must use either the theme or plugin option. The latter is recommended.
//
//## OPTIONS
//
//<slug>
//: The internal name of the block.
//
//[--title=<title>]
//: The display title for your block.
//
//[--dashicon=<dashicon>]
//: The dashicon to make it easier to identify your block.
//
//[--category=<category>]
//: The category name to help users browse and discover your block.
//---
//default: widgets
//options:
//  - common
//  - embed
//  - formatting
//  - layout
//  - widgets
//---
//
//[--theme]
//: Create files in the active theme directory. Specify a theme with `--theme=<theme>` to have the file placed in that theme.
//
//[--plugin=<plugin>]
//: Create files in the given plugin's directory.
//
//[--force]
//: Overwrite files that already exist.
//
//## EXAMPLES
//
//    # Generate a 'movie' block for the 'movies' plugin
//    $ wp scaffold block movie --title="Movie block" --plugin=movies
//    Success: Created block 'Movie block'.
//
//    # Generate a 'movie' block for the 'simple-life' theme
//    $ wp scaffold block movie --title="Movie block" --theme=simple-life
//     Success: Created block 'Movie block'.
//
//    # Create a new plugin and add two blocks
//    # Create plugin called books
//    $ wp scaffold plugin books
//    # Add a block called book to plugin books
//    $ wp scaffold block book --title="Book" --plugin=books
//    # Add a second block to plugin called books.
//    $ wp scaffold block books --title="Book List" --plugin=books
//
//