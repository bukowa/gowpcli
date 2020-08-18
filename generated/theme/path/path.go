package path

//Gets the path to a theme or to the theme directory.

type Path struct {
    
    Theme string // [<theme>]
    
    Dir bool // [--dir]
    
}

//## OPTIONS
//
//[<theme>]
//: The theme to get the path to. Path includes "style.css" file.
//If not set, will return the path to the themes directory.
//
//[--dir]
//: If set, get the path to the closest parent directory, instead of the
//theme's "style.css" file.
//
//## EXAMPLES
//
//    # Get theme path
//    $ wp theme path
//    /var/www/example.com/public_html/wp-content/themes
//
//    # Change directory to theme path
//    $ cd $(wp theme path)
//
//