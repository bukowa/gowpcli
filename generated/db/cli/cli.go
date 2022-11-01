/*
## OPTIONS
	[--database=<database>]
	: Use a specific database. Defaults to DB_NAME.
	[--default-character-set=<character-set>]
	: Use a specific character set. Defaults to DB_CHARSET when defined.
	[--dbuser=<value>]
	: Username to pass to mysql. Defaults to DB_USER.
	[--dbpass=<value>]
	: Password to pass to mysql. Defaults to DB_PASSWORD.
	[--<field>=<value>]
	: Extra arguments to pass to mysql. [Refer to mysql docs](https://dev.mysql.com/doc/en/mysql-command-options.html).
	[--defaults]
	: Loads the environment's MySQL option files. Default behavior is to skip loading them to avoid failures due to misconfiguration.
## EXAMPLES
	    # Open MySQL console
	    $ wp db cli
	    mysql>
	
 */
package cli
import utils "github.com/bukowa/gowpcli"

// Cli //Opens a MySQL console using credentials from wp-config.php
type Cli struct {
    Database string // [--database=<database>]
    DefaultCharacterSet string // [--default-character-set=<character-set>]
    Dbuser string // [--dbuser=<value>]
    Dbpass string // [--dbpass=<value>]
    FieldMap map[string]string // [--<field>=<value>]
    Defaults bool // [--defaults]
}

func (c Cli) Args() []string {
    var args = []string{"db", "cli"}
    args = utils.MakeArg(args, "[--database=<database>]", c.Database)
    args = utils.MakeArg(args, "[--default-character-set=<character-set>]", c.DefaultCharacterSet)
    args = utils.MakeArg(args, "[--dbuser=<value>]", c.Dbuser)
    args = utils.MakeArg(args, "[--dbpass=<value>]", c.Dbpass)
    args = utils.MakeArg(args, "[--<field>=<value>]", c.FieldMap)
    args = utils.MakeArg(args, "[--defaults]", c.Defaults)
    return args
}

