package create

//Generates a wp-config.php file.

type Create struct {
    
    Dbname string // --dbname=<dbname>
    
    Dbuser string // --dbuser=<dbuser>
    
    Dbpass string // [--dbpass=<dbpass>]
    
    Dbhost string // [--dbhost=<dbhost>]
    
    Dbprefix string // [--dbprefix=<dbprefix>]
    
    Dbcharset string // [--dbcharset=<dbcharset>]
    
    Dbcollate string // [--dbcollate=<dbcollate>]
    
    Locale string // [--locale=<locale>]
    
    ExtraPhp bool // [--extra-php]
    
    SkipSalts bool // [--skip-salts]
    
    SkipCheck bool // [--skip-check]
    
    Force bool // [--force]
    
}

//Creates a new wp-config.php with database constants, and verifies that
//the database constants are correct.
//
//## OPTIONS
//
//--dbname=<dbname>
//: Set the database name.
//
//--dbuser=<dbuser>
//: Set the database user.
//
//[--dbpass=<dbpass>]
//: Set the database user password.
//
//[--dbhost=<dbhost>]
//: Set the database host.
//---
//default: localhost
//---
//
//[--dbprefix=<dbprefix>]
//: Set the database table prefix.
//---
//default: wp_
//---
//
//[--dbcharset=<dbcharset>]
//: Set the database charset.
//---
//default: utf8
//---
//
//[--dbcollate=<dbcollate>]
//: Set the database collation.
//---
//default:
//---
//
//[--locale=<locale>]
//: Set the WPLANG constant. Defaults to $wp_local_package variable.
//
//[--extra-php]
//: If set, the command copies additional PHP code into wp-config.php from STDIN.
//
//[--skip-salts]
//: If set, keys and salts won't be generated, but should instead be passed via `--extra-php`.
//
//[--skip-check]
//: If set, the database connection is not checked.
//
//[--force]
//: Overwrites existing files, if present.
//
//## EXAMPLES
//
//    # Standard wp-config.php file
//    $ wp config create --dbname=testing --dbuser=wp --dbpass=securepswd --locale=ro_RO
//    Success: Generated 'wp-config.php' file.
//
//    # Enable WP_DEBUG and WP_DEBUG_LOG
//    $ wp config create --dbname=testing --dbuser=wp --dbpass=securepswd --extra-php <<PHP
//    define( 'WP_DEBUG', true );
//    define( 'WP_DEBUG_LOG', true );
//    PHP
//    Success: Generated 'wp-config.php' file.
//
//    # Avoid disclosing password to bash history by reading from password.txt
//    # Using --prompt=dbpass will prompt for the 'dbpass' argument
//    $ wp config create --dbname=testing --dbuser=wp --prompt=dbpass < password.txt
//    Success: Generated 'wp-config.php' file.
//
//