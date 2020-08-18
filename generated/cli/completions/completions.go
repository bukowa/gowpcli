package completions

//Generates tab completion strings.

type Completions struct {
    
    Line string // --line=<line>
    
    Point string // --point=<point>
    
}

//## OPTIONS
//
//--line=<line>
//: The current command line to be executed.
//
//--point=<point>
//: The index to the current cursor position relative to the beginning of the command.
//
//## EXAMPLES
//
//    # Generate tab completion strings.
//    $ wp cli completions --line='wp eva' --point=100
//    eval
//    eval-file
//
//