package activate

//Activates a theme.

type Activate struct {
    
    Theme string // <theme>
    
}

//## OPTIONS
//
//<theme>
//: The theme to activate.
//
//## EXAMPLES
//
//    $ wp theme activate twentysixteen
//    Success: Switched to 'Twenty Sixteen' theme.
//
//