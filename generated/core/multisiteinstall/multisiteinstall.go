package multisiteinstall

//Installs WordPress multisite from scratch.

type MultisiteInstall struct {
    
    Url string // [--url=<url>]
    
    Base string // [--base=<url-path>]
    
    Subdomains bool // [--subdomains]
    
    Title string // --title=<site-title>
    
    AdminUser string // --admin_user=<username>
    
    AdminPassword string // [--admin_password=<password>]
    
    AdminEmail string // --admin_email=<email>
    
    SkipEmail bool // [--skip-email]
    
    SkipConfig bool // [--skip-config]
    
}

//Creates the WordPress tables in the database using the URL, title, and
//default admin user details provided. Then, creates the multisite tables
//in the database and adds multisite constants to the wp-config.php.
//
//For those using WordPress with Apache, remember to update the `.htaccess`
//file with the appropriate multisite rewrite rules.
//
//## OPTIONS
//
//[--url=<url>]
//: The address of the new site.
//
//[--base=<url-path>]
//: Base path after the domain name that each site url in the network will start with.
//---
//default: /
//---
//
//[--subdomains]
//: If passed, the network will use subdomains, instead of subdirectories. Doesn't work with 'localhost'.
//
//--title=<site-title>
//: The title of the new site.
//
//--admin_user=<username>
//: The name of the admin user.
//---
//default: admin
//---
//
//[--admin_password=<password>]
//: The password for the admin user. Defaults to randomly generated string.
//
//--admin_email=<email>
//: The email address for the admin user.
//
//[--skip-email]
//: Don't send an email notification to the new admin user.
//
//[--skip-config]
//: Don't add multisite constants to wp-config.php.
//
//## EXAMPLES
//
//    $ wp core multisite-install --title="Welcome to the WordPress" \
//    > --admin_user="admin" --admin_password="password" \
//    > --admin_email="user@example.com"
//    Single site database tables already present.
//    Set up multisite database tables.
//    Added multisite constants to wp-config.php.
//    Success: Network installed. Don't forget to set up rewrite rules.
//
//