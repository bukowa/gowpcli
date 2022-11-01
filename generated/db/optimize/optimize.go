/*
## INFO
	Runs `mysqlcheck` utility with `--optimize=true` using `DB_HOST`,
	`DB_NAME`, `DB_USER` and `DB_PASSWORD` database credentials
	specified in wp-config.php.
	[See docs](http://dev.mysql.com/doc/refman/5.7/en/optimize-table.html)
	for more details on the `OPTIMIZE TABLE` statement.
## OPTIONS
	[--dbuser=<value>]
	: Username to pass to mysqlcheck. Defaults to DB_USER.
	[--dbpass=<value>]
	: Password to pass to mysqlcheck. Defaults to DB_PASSWORD.
	[--<field>=<value>]
	: Extra arguments to pass to mysqlcheck. [Refer to mysqlcheck docs](https://dev.mysql.com/doc/en/mysqlcheck.html).
	[--defaults]
	: Loads the environment's MySQL option files. Default behavior is to skip loading them to avoid failures due to misconfiguration.
## EXAMPLES
	    $ wp db optimize
	    Success: Database optimized.
	
 */
package optimize
import utils "github.com/bukowa/gowpcli"

// Optimize //Optimizes the database.
type Optimize struct {
    Dbuser string // [--dbuser=<value>]
    Dbpass string // [--dbpass=<value>]
    FieldMap map[string]string // [--<field>=<value>]
    Defaults bool // [--defaults]
}

func (o Optimize) Args() []string {
    var args = []string{"db", "optimize"}
    args = utils.MakeArg(args, "[--dbuser=<value>]", o.Dbuser)
    args = utils.MakeArg(args, "[--dbpass=<value>]", o.Dbpass)
    args = utils.MakeArg(args, "[--<field>=<value>]", o.FieldMap)
    args = utils.MakeArg(args, "[--defaults]", o.Defaults)
    return args
}

