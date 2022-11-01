/*
## INFO
	Runs `DROP_DATABASE` and `CREATE_DATABASE` SQL statements using
	`DB_HOST`, `DB_NAME`, `DB_USER` and `DB_PASSWORD` database credentials
	specified in wp-config.php.
## OPTIONS
	[--dbuser=<value>]
	: Username to pass to mysql. Defaults to DB_USER.
	[--dbpass=<value>]
	: Password to pass to mysql. Defaults to DB_PASSWORD.
	[--yes]
	: Answer yes to the confirmation message.
	[--defaults]
	: Loads the environment's MySQL option files. Default behavior is to skip loading them to avoid failures due to misconfiguration.
## EXAMPLES
	    $ wp db reset --yes
	    Success: Database reset.
	
 */
package reset
import utils "github.com/bukowa/gowpcli"

// Reset //Removes all tables from the database.
type Reset struct {
    Dbuser string // [--dbuser=<value>]
    Dbpass string // [--dbpass=<value>]
    Yes bool // [--yes]
    Defaults bool // [--defaults]
}

func (r Reset) Args() []string {
    var args = []string{"db", "reset"}
    args = utils.MakeArg(args, "[--dbuser=<value>]", r.Dbuser)
    args = utils.MakeArg(args, "[--dbpass=<value>]", r.Dbpass)
    args = utils.MakeArg(args, "[--yes]", r.Yes)
    args = utils.MakeArg(args, "[--defaults]", r.Defaults)
    return args
}

