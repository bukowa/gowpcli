package repair

//Repairs the database.

type Repair struct {
    
    Dbuser string // [--dbuser=<value>]
    
    Dbpass string // [--dbpass=<value>]
    
    FieldMap map[string]string // [--<field>=<value>]
    
}

//Runs `mysqlcheck` utility with `--repair=true` using `DB_HOST`,
//`DB_NAME`, `DB_USER` and `DB_PASSWORD` database credentials
//specified in wp-config.php.
//
//[See docs](http://dev.mysql.com/doc/refman/5.7/en/repair-table.html) for
//more details on the `REPAIR TABLE` statement.
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
//    $ wp db repair
//    Success: Database repaired.
//
//