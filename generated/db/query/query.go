/*
## INFO
	Executes an arbitrary SQL query using `DB_HOST`, `DB_NAME`, `DB_USER`
	 and `DB_PASSWORD` database credentials specified in wp-config.php.
## OPTIONS
	[<sql>]
	: A SQL query. If not passed, will try to read from STDIN.
	[--dbuser=<value>]
	: Username to pass to mysql. Defaults to DB_USER.
	[--dbpass=<value>]
	: Password to pass to mysql. Defaults to DB_PASSWORD.
	[--<field>=<value>]
	: Extra arguments to pass to mysql. [Refer to mysql docs](https://dev.mysql.com/doc/en/mysql-command-options.html).
	[--defaults]
	: Loads the environment's MySQL option files. Default behavior is to skip loading them to avoid failures due to misconfiguration.
## EXAMPLES
	    # Execute a query stored in a file
	    $ wp db query < debug.sql
	    # Check all tables in the database
	    $ wp db query "CHECK TABLE $(wp db tables | paste -s -d, -);"
	    +---------------------------------------+-------+----------+----------+
	    | Table                                 | Op    | Msg_type | Msg_text |
	    +---------------------------------------+-------+----------+----------+
	    | wordpress_dbase.wp_users              | check | status   | OK       |
	    | wordpress_dbase.wp_usermeta           | check | status   | OK       |
	    | wordpress_dbase.wp_posts              | check | status   | OK       |
	    | wordpress_dbase.wp_comments           | check | status   | OK       |
	    | wordpress_dbase.wp_links              | check | status   | OK       |
	    | wordpress_dbase.wp_options            | check | status   | OK       |
	    | wordpress_dbase.wp_postmeta           | check | status   | OK       |
	    | wordpress_dbase.wp_terms              | check | status   | OK       |
	    | wordpress_dbase.wp_term_taxonomy      | check | status   | OK       |
	    | wordpress_dbase.wp_term_relationships | check | status   | OK       |
	    | wordpress_dbase.wp_termmeta           | check | status   | OK       |
	    | wordpress_dbase.wp_commentmeta        | check | status   | OK       |
	    +---------------------------------------+-------+----------+----------+
	    # Pass extra arguments through to MySQL
	    $ wp db query 'SELECT * FROM wp_options WHERE option_name="home"' --skip-column-names
	    +---+------+------------------------------+-----+
	    | 2 | home | http://wordpress-develop.dev | yes |
	    +---+------+------------------------------+-----+
	
 */
package query
import utils "github.com/bukowa/gowpcli"

// Query //Executes a SQL query against the database.
type Query struct {
    Sql string // [<sql>]
    Dbuser string // [--dbuser=<value>]
    Dbpass string // [--dbpass=<value>]
    FieldMap map[string]string // [--<field>=<value>]
    Defaults bool // [--defaults]
}

func (q Query) Args() []string {
    var args = []string{"db", "query"}
    args = utils.MakeArg(args, "[<sql>]", q.Sql)
    args = utils.MakeArg(args, "[--dbuser=<value>]", q.Dbuser)
    args = utils.MakeArg(args, "[--dbpass=<value>]", q.Dbpass)
    args = utils.MakeArg(args, "[--<field>=<value>]", q.FieldMap)
    args = utils.MakeArg(args, "[--defaults]", q.Defaults)
    return args
}

