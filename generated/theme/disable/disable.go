package disable

//Disables a theme on a WordPress multisite install.

type Disable struct {
    
    Theme string // <theme>
    
    Network bool // [--network]
    
}

//Removes ability for a theme to be activated from the dashboard of a site
//on a WordPress multisite install.
//
//## OPTIONS
//
//<theme>
//: The theme to disable.
//
//[--network]
//: If set, the theme is disabled on the network level. Note that
//individual sites may still have this theme enabled if it was
//enabled for them independently.
//
//## EXAMPLES
//
//    # Disable theme
//    $ wp theme disable twentysixteen
//    Success: Disabled the 'Twenty Sixteen' theme.
//
//    # Disable theme in network level
//    $ wp theme disable twentysixteen --network
//    Success: Network disabled the 'Twenty Sixteen' theme.
//
//