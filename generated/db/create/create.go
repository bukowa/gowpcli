package create

//Creates a new database.

type Create struct {
    
    Dbuser string // [--dbuser=<value>]
    
    Dbpass string // [--dbpass=<value>]
    
}

//Runs `CREATE_DATABASE` SQL statement using `DB_HOST`, `DB_NAME`,
//`DB_USER` and `DB_PASSWORD` database credentials specified in
//wp-config.php.
//
//## OPTIONS
//
//[--dbuser=<value>]
//: Username to pass to mysql. Defaults to DB_USER.
//
//[--dbpass=<value>]
//: Password to pass to mysql. Defaults to DB_PASSWORD.
//
//## EXAMPLES
//
//    $ wp db create
//    Success: Database created.
//
//