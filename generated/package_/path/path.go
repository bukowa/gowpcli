package path

//Gets the path to an installed WP-CLI package, or the package directory.

type Path struct {
    
    Name string // [<name>]
    
}

//If you want to contribute to a package, this is a great way to jump to it.
//
//## OPTIONS
//
//[<name>]
//: Name of the package to get the directory for.
//
//## EXAMPLES
//
//    # Get package path
//    $ wp package path
//    /home/person/.wp-cli/packages/
//
//    # Change directory to package path
//    $ cd $(wp package path) && pwd
//    /home/vagrant/.wp-cli/packages
//
//