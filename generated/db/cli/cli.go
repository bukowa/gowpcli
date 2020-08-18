package cli

//Opens a MySQL console using credentials from wp-config.php

type Cli struct {
    
    Database string // [--database=<database>]
    
    DefaultCharacterSet string // [--default-character-set=<character-set>]
    
    Dbuser string // [--dbuser=<value>]
    
    Dbpass string // [--dbpass=<value>]
    
    FieldMap map[string]string // [--<field>=<value>]
    
}

//## OPTIONS
//
//[--database=<database>]
//: Use a specific database. Defaults to DB_NAME.
//
//[--default-character-set=<character-set>]
//: Use a specific character set. Defaults to DB_CHARSET when defined.
//
//[--dbuser=<value>]
//: Username to pass to mysql. Defaults to DB_USER.
//
//[--dbpass=<value>]
//: Password to pass to mysql. Defaults to DB_PASSWORD.
//
//[--<field>=<value>]
//: Extra arguments to pass to mysql. [Refer to mysql docs](https://dev.mysql.com/doc/en/mysql-command-options.html).
//
//## EXAMPLES
//
//    # Open MySQL console
//    $ wp db cli
//    mysql>
//
//