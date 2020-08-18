/*
## INFO
	Runs `DROP_DATABASE` SQL statement using `DB_HOST`, `DB_NAME`,
	`DB_USER` and `DB_PASSWORD` database credentials specified in
	wp-config.php.
## OPTIONS
	[--dbuser=<value>]
	: Username to pass to mysql. Defaults to DB_USER.
	[--dbpass=<value>]
	: Password to pass to mysql. Defaults to DB_PASSWORD.
	[--yes]
	: Answer yes to the confirmation message.
## EXAMPLES
	    $ wp db drop --yes
	    Success: Database dropped.
	
 */
package drop
import utils "github.com/bukowa/gowpcli"

// Drop //Deletes the existing database.
type Drop struct {
    Dbuser string // [--dbuser=<value>]
    Dbpass string // [--dbpass=<value>]
    Yes bool // [--yes]
}

func (d Drop) Args() []string {
    var args = []string{"db", "drop"}
    args = utils.MakeArg(args, "[--dbuser=<value>]", d.Dbuser)
    args = utils.MakeArg(args, "[--dbpass=<value>]", d.Dbpass)
    args = utils.MakeArg(args, "[--yes]", d.Yes)
    return args
}

