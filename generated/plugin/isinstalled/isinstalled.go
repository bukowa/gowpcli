package isinstalled

//Checks if a given plugin is installed.

type IsInstalled struct {
    
    Plugin string // <plugin>
    
}

//Returns exit code 0 when installed, 1 when uninstalled.
//
//## OPTIONS
//
//<plugin>
//: The plugin to check.
//
//## EXAMPLES
//
//    # Check whether plugin is installed; exit status 0 if installed, otherwise 1
//    $ wp plugin is-installed hello
//    $ echo $?
//    1
//
//