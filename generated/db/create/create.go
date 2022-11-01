/*
## INFO
	Runs `CREATE_DATABASE` SQL statement using `DB_HOST`, `DB_NAME`,
	`DB_USER` and `DB_PASSWORD` database credentials specified in
	wp-config.php.
## OPTIONS
	[--dbuser=<value>]
	: Username to pass to mysql. Defaults to DB_USER.
	[--dbpass=<value>]
	: Password to pass to mysql. Defaults to DB_PASSWORD.
	[--defaults]
	: Loads the environment's MySQL option files. Default behavior is to skip loading them to avoid failures due to misconfiguration.
## EXAMPLES
	    $ wp db create
	    Success: Database created.
	
 */
package create
import utils "github.com/bukowa/gowpcli"

// Create //Creates a new database.
type Create struct {
    Dbuser string // [--dbuser=<value>]
    Dbpass string // [--dbpass=<value>]
    Defaults bool // [--defaults]
}

func (c Create) Args() []string {
    var args = []string{"db", "create"}
    args = utils.MakeArg(args, "[--dbuser=<value>]", c.Dbuser)
    args = utils.MakeArg(args, "[--dbpass=<value>]", c.Dbpass)
    args = utils.MakeArg(args, "[--defaults]", c.Defaults)
    return args
}

