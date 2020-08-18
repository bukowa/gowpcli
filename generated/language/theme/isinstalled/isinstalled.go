package isinstalled

//Checks if a given language is installed.

type IsInstalled struct {
    
    Theme string // <theme>
    
    Language []string // <language>...
    
}

//Returns exit code 0 when installed, 1 when uninstalled.
//
//## OPTIONS
//
//<theme>
//: Theme to check for.
//
//<language>...
//: The language code to check.
//
//## EXAMPLES
//
//    # Check whether the German language is installed for Twenty Seventeen; exit status 0 if installed, otherwise 1.
//    $ wp language theme is-installed twentyseventeen de_DE
//    $ echo $?
//    1
//
//