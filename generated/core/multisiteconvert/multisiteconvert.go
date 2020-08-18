package multisiteconvert

//Transforms an existing single-site installation into a multisite installation.

type MultisiteConvert struct {
    
    Title string // [--title=<network-title>]
    
    Base string // [--base=<url-path>]
    
    Subdomains bool // [--subdomains]
    
}

//Creates the multisite database tables, and adds the multisite constants
//to wp-config.php.
//
//For those using WordPress with Apache, remember to update the `.htaccess`
//file with the appropriate multisite rewrite rules.
//
//[Review the multisite documentation](https://codex.wordpress.org/Create_A_Network)
//for more details about how multisite works.
//
//## OPTIONS
//
//[--title=<network-title>]
//: The title of the new network.
//
//[--base=<url-path>]
//: Base path after the domain name that each site url will start with.
//---
//default: /
//---
//
//[--subdomains]
//: If passed, the network will use subdomains, instead of subdirectories. Doesn't work with 'localhost'.
//
//## EXAMPLES
//
//    $ wp core multisite-convert
//    Set up multisite database tables.
//    Added multisite constants to wp-config.php.
//    Success: Network installed. Don't forget to set up rewrite rules.
//
//