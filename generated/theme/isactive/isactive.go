package isactive

//Checks if a given theme is active.

type IsActive struct {
    
    Theme string // <theme>
    
}

//Returns exit code 0 when active, 1 when not active.
//
//## OPTIONS
//
//<theme>
//: The plugin to check.
//
//## EXAMPLES
//
//    # Check whether theme is Active; exit status 0 if active, otherwise 1
//    $ wp theme is-active twentyfifteen
//    $ echo $?
//    1
//
//