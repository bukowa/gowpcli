package clean

//Removes all tables with `$table_prefix` from the database.

type Clean struct {
    
    Dbuser string // [--dbuser=<value>]
    
    Dbpass string // [--dbpass=<value>]
    
    Yes bool // [--yes]
    
}

//Runs `DROP_TABLE` for each table that has a `$table_prefix` as specified
//in wp-config.php.
//
//## OPTIONS
//
//[--dbuser=<value>]
//: Username to pass to mysql. Defaults to DB_USER.
//
//[--dbpass=<value>]
//: Password to pass to mysql. Defaults to DB_PASSWORD.
//
//[--yes]
//: Answer yes to the confirmation message.
//
//## EXAMPLES
//
//    # Delete all tables that match the current site prefix.
//    $ wp db clean --yes
//    Success: Tables dropped.
//
//