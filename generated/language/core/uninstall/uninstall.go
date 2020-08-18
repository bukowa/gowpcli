package uninstall

//Uninstalls a given language.

type Uninstall struct {
    
    Language []string // <language>...
    
}

//## OPTIONS
//
//<language>...
//: Language code to uninstall.
//
//## EXAMPLES
//
//    $ wp language core uninstall ja
//    Success: Language uninstalled.
//
//