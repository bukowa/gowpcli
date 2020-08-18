package uninstall

//Uninstalls a WP-CLI package.

type Uninstall struct {
    
    Name string // <name>
    
}

//## OPTIONS
//
//<name>
//: Name of the package to uninstall.
//
//## EXAMPLES
//
//    $ wp package uninstall wp-cli/server-command
//    Removing require statement from /home/person/.wp-cli/packages/composer.json
//    Deleting package directory /home/person/.wp-cli/packages/vendor/wp-cli/server-command
//    Regenerating Composer autoload.
//    Success: Uninstalled package.
//
//