package isinstalled

//Checks if a given language is installed.

type IsInstalled struct {
    
    Language string // <language>
    
}

//Returns exit code 0 when installed, 1 when uninstalled.
//
//## OPTIONS
//
//<language>
//: The language code to check.
//
//## EXAMPLES
//
//    # Check whether the German language is installed; exit status 0 if installed, otherwise 1.
//    $ wp language core is-installed de_DE
//    $ echo $?
//    1
//
//