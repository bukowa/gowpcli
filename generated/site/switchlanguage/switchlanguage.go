package switchlanguage

//Activates a given language.

type SwitchLanguage struct {
    
    Language string // <language>
    
}

//## OPTIONS
//
//<language>
//: Language code to activate.
//
//## EXAMPLES
//
//    $ wp site switch-language ja
//    Success: Language activated.
//
//