package shell

//Opens an interactive PHP console for running and testing PHP code.

type Shell struct {
    
    Basic bool // [--basic]
    
}

//`wp shell` allows you to evaluate PHP statements and expressions
//interactively, from within a WordPress environment. Type a bit of code,
//hit enter, and see the code execute right before you. Because WordPress
//is loaded, you have access to all the functions, classes and globals
//that you can use within a WordPress plugin, for example.
//
//## OPTIONS
//
//[--basic]
//: Force the use of WP-CLI's built-in PHP REPL, even if the Boris or
//PsySH PHP REPLs are available.
//
//## EXAMPLES
//
//    # Call get_bloginfo() to get the name of the site.
//    $ wp shell
//    wp> get_bloginfo( 'name' );
//    => string(6) "WP-CLI"
//
//