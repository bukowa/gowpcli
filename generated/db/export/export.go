package export

//Exports the database to a file or to STDOUT.

type Export struct {
    
    File string // [<file>]
    
    Dbuser string // [--dbuser=<value>]
    
    Dbpass string // [--dbpass=<value>]
    
    FieldMap map[string]string // [--<field>=<value>]
    
    Tables string // [--tables=<tables>]
    
    ExcludeTables string // [--exclude_tables=<tables>]
    
    Porcelain bool // [--porcelain]
    
}

//Runs `mysqldump` utility using `DB_HOST`, `DB_NAME`, `DB_USER` and
//`DB_PASSWORD` database credentials specified in wp-config.php.
//
//## OPTIONS
//
//[<file>]
//: The name of the SQL file to export. If '-', then outputs to STDOUT. If
//omitted, it will be '{dbname}-{Y-m-d}-{random-hash}.sql'.
//
//[--dbuser=<value>]
//: Username to pass to mysqldump. Defaults to DB_USER.
//
//[--dbpass=<value>]
//: Password to pass to mysqldump. Defaults to DB_PASSWORD.
//
//[--<field>=<value>]
//: Extra arguments to pass to mysqldump. [Refer to mysqldump docs](https://dev.mysql.com/doc/en/mysqldump.html#mysqldump-option-summary).
//
//[--tables=<tables>]
//: The comma separated list of specific tables to export. Excluding this parameter will export all tables in the database.
//
//[--exclude_tables=<tables>]
//: The comma separated list of specific tables that should be skipped from exporting. Excluding this parameter will export all tables in the database.
//
//[--porcelain]
//: Output filename for the exported database.
//
//## EXAMPLES
//
//    # Export database with drop query included
//    $ wp db export --add-drop-table
//    Success: Exported to 'wordpress_dbase-db72bb5.sql'.
//
//    # Export certain tables
//    $ wp db export --tables=wp_options,wp_users
//    Success: Exported to 'wordpress_dbase-db72bb5.sql'.
//
//    # Export all tables matching a wildcard
//    $ wp db export --tables=$(wp db tables 'wp_user*' --format=csv)
//    Success: Exported to 'wordpress_dbase-db72bb5.sql'.
//
//    # Export all tables matching prefix
//    $ wp db export --tables=$(wp db tables --all-tables-with-prefix --format=csv)
//    Success: Exported to 'wordpress_dbase-db72bb5.sql'.
//
//    # Export certain posts without create table statements
//    $ wp db export --no-create-info=true --tables=wp_posts --where="ID in (100,101,102)"
//    Success: Exported to 'wordpress_dbase-db72bb5.sql'.
//
//    # Export relating meta for certain posts without create table statements
//    $ wp db export --no-create-info=true --tables=wp_postmeta --where="post_id in (100,101,102)"
//    Success: Exported to 'wordpress_dbase-db72bb5.sql'.
//
//    # Skip certain tables from the exported database
//    $ wp db export --exclude_tables=wp_options,wp_users
//    Success: Exported to 'wordpress_dbase-db72bb5.sql'.
//
//    # Skip all tables matching a wildcard from the exported database
//    $ wp db export --exclude_tables=$(wp db tables 'wp_user*' --format=csv)
//    Success: Exported to 'wordpress_dbase-db72bb5.sql'.
//
//    # Skip all tables matching prefix from the exported database
//    $ wp db export --exclude_tables=$(wp db tables --all-tables-with-prefix --format=csv)
//    Success: Exported to 'wordpress_dbase-db72bb5.sql'.
//
//    # Export database to STDOUT.
//    $ wp db export -
//    -- MySQL dump 10.13  Distrib 5.7.19, for osx10.12 (x86_64)
//    --
//    -- Host: localhost    Database: wpdev
//    -- ------------------------------------------------------
//    -- Server version    5.7.19
//    ...
//
//