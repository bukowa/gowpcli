package evalfile

//Loads and executes a PHP file.

type EvalFile struct {
    
    File string // <file>
    
    Arg []string // [<arg>...]
    
    SkipWordpress bool // [--skip-wordpress]
    
}

//Note: because code is executed within a method, global variables need
//to be explicitly globalized.
//
//## OPTIONS
//
//<file>
//: The path to the PHP file to execute.  Use '-' to run code from STDIN.
//
//[<arg>...]
//: One or more arguments to pass to the file. They are placed in the $args variable.
//
//[--skip-wordpress]
//: Load and execute file without loading WordPress.
//
//