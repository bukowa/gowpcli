package deactivate

//Deactivates one or more plugins.

type Deactivate struct {
    
    Plugin []string // [<plugin>...]
    
    Uninstall bool // [--uninstall]
    
    All bool // [--all]
    
    Network bool // [--network]
    
}

//## OPTIONS
//
//[<plugin>...]
//: One or more plugins to deactivate.
//
//[--uninstall]
//: Uninstall the plugin after deactivation.
//
//[--all]
//: If set, all plugins will be deactivated.
//
//[--network]
//: If set, the plugin will be deactivated for the entire multisite network.
//
//## EXAMPLES
//
//    # Deactivate plugin
//    $ wp plugin deactivate hello
//    Plugin 'hello' deactivated.
//    Success: Deactivated 1 of 1 plugins.
//
//