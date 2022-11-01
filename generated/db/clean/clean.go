/*
## INFO
	Runs `DROP_TABLE` for each table that has a `$table_prefix` as specified
	in wp-config.php.
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
	    # Delete all tables that match the current site prefix.
	    $ wp db clean --yes
	    Success: Tables dropped.
	
 */
package clean
import utils "github.com/bukowa/gowpcli"

// Clean //Removes all tables with `$table_prefix` from the database.
type Clean struct {
    Dbuser string // [--dbuser=<value>]
    Dbpass string // [--dbpass=<value>]
    Yes bool // [--yes]
    Defaults bool // [--defaults]
}

func (c Clean) Args() []string {
    var args = []string{"db", "clean"}
    args = utils.MakeArg(args, "[--dbuser=<value>]", c.Dbuser)
    args = utils.MakeArg(args, "[--dbpass=<value>]", c.Dbpass)
    args = utils.MakeArg(args, "[--yes]", c.Yes)
    args = utils.MakeArg(args, "[--defaults]", c.Defaults)
    return args
}

