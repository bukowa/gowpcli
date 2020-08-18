package check

//Checks the current status of the database.

type Check struct {
    
    Dbuser string // [--dbuser=<value>]
    
    Dbpass string // [--dbpass=<value>]
    
    FieldMap map[string]string // [--<field>=<value>]
    
}

//Runs `mysqlcheck` utility with `--check` using `DB_HOST`,
//`DB_NAME`, `DB_USER` and `DB_PASSWORD` database credentials
//specified in wp-config.php.
//
//[See docs](http://dev.mysql.com/doc/refman/5.7/en/check-table.html)
//for more details on the `CHECK TABLE` statement.
//
//## OPTIONS
//
//[--dbuser=<value>]
//: Username to pass to mysqlcheck. Defaults to DB_USER.
//
//[--dbpass=<value>]
//: Password to pass to mysqlcheck. Defaults to DB_PASSWORD.
//
//[--<field>=<value>]
//: Extra arguments to pass to mysqlcheck. [Refer to mysqlcheck docs](https://dev.mysql.com/doc/en/mysqlcheck.html).
//
//## EXAMPLES
//
//    $ wp db check
//    Success: Database checked.
//
//