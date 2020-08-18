package activate

//Activates a given language.

type Activate struct {
    
    Language string // <language>
    
}

//## OPTIONS
//
//<language>
//: Language code to activate.
//
//## EXAMPLES
//
//    $ wp language core activate ja
//    Success: Language activated.
//
//