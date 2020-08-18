package reset

//Removes all tables from the database.

type Reset struct {
    
    Dbuser string // [--dbuser=<value>]
    
    Dbpass string // [--dbpass=<value>]
    
    Yes bool // [--yes]
    
}

//Runs `DROP_DATABASE` and `CREATE_DATABASE` SQL statements using
//`DB_HOST`, `DB_NAME`, `DB_USER` and `DB_PASSWORD` database credentials
//specified in wp-config.php.
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
//    $ wp db reset --yes
//    Success: Database reset.
//
//