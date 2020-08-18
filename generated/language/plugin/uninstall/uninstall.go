package uninstall

//Uninstalls a given language for a plugin.

type Uninstall struct {
    
    Plugin string // <plugin>
    
    Language []string // <language>...
    
}

//## OPTIONS
//
//<plugin>
//: Plugin to uninstall language for.
//
//<language>...
//: Language code to uninstall.
//
//## EXAMPLES
//
//    $ wp language plugin uninstall hello-dolly ja
//    Success: Language uninstalled.
//
//