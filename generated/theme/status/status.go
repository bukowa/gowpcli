package status

//Reveals the status of one or all themes.

type Status struct {
    
    Theme string // [<theme>]
    
}

//## OPTIONS
//
//[<theme>]
//: A particular theme to show the status for.
//
//## EXAMPLES
//
//    $ wp theme status twentysixteen
//    Theme twentysixteen details:
//         Name: Twenty Sixteen
//         Status: Inactive
//         Version: 1.2
//         Author: the WordPress team
//
//