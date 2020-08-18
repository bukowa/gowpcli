package activate

//Activates one or more plugins.

type Activate struct {
    
    Plugin []string // [<plugin>...]
    
    All bool // [--all]
    
    Network bool // [--network]
    
}

//## OPTIONS
//
//[<plugin>...]
//: One or more plugins to activate.
//
//[--all]
//: If set, all plugins will be activated.
//
//[--network]
//: If set, the plugin will be activated for the entire multisite network.
//
//## EXAMPLES
//
//    # Activate plugin
//    $ wp plugin activate hello
//    Plugin 'hello' activated.
//    Success: Activated 1 of 1 plugins.
//
//    # Activate plugin in entire multisite network
//    $ wp plugin activate hello --network
//    Plugin 'hello' network activated.
//    Success: Network activated 1 of 1 plugins.
//
//