package uninstall

//Uninstalls a given language for a theme.

type Uninstall struct {
    
    Theme string // <theme>
    
    Language []string // <language>...
    
}

//## OPTIONS
//
//<theme>
//: Theme to uninstall language for.
//
//<language>...
//: Language code to uninstall.
//
//## EXAMPLES
//
//    $ wp language theme uninstall twentyten ja
//    Success: Language uninstalled.
//
//