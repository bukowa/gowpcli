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
}

func (c Create) Args() []string {
    var args = []string{"db", "create"}
    args = utils.MakeArg(args, "[--dbuser=<value>]", c.Dbuser)
    args = utils.MakeArg(args, "[--dbpass=<value>]", c.Dbpass)
    return args
}

