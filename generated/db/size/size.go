/*
## INFO
	Display the database name and size for `DB_NAME` specified in wp-config.php.
	The size defaults to a human-readable number.
	Available size formats include:
	* b (bytes)
	* kb (kilobytes)
	* mb (megabytes)
	* gb (gigabytes)
	* tb (terabytes)
	* B   (ISO Byte setting, with no conversion)
	* KB  (ISO Kilobyte setting, with 1 KB  = 1,000 B)
	* KiB (ISO Kibibyte setting, with 1 KiB = 1,024 B)
	* MB  (ISO Megabyte setting, with 1 MB  = 1,000 KB)
	* MiB (ISO Mebibyte setting, with 1 MiB = 1,024 KiB)
	* GB  (ISO Gigabyte setting, with 1 GB  = 1,000 MB)
	* GiB (ISO Gibibyte setting, with 1 GiB = 1,024 MiB)
	* TB  (ISO Terabyte setting, with 1 TB  = 1,000 GB)
	* TiB (ISO Tebibyte setting, with 1 TiB = 1,024 GiB)
## OPTIONS
	[--size_format=<format>]
	: Display the database size only, as a bare number.
	---
	options:
	  - b
	  - kb
	  - mb
	  - gb
	  - tb
	  - B
	  - KB
	  - KiB
	  - MB
	  - MiB
	  - GB
	  - GiB
	  - TB
	  - TiB
	---
	[--tables]
	: Display each table name and size instead of the database size.
	[--human-readable]
	: Display database sizes in human readable formats.
	[--format=<format>]
	: Render output in a particular format.
	---
	options:
	  - table
	  - csv
	  - json
	  - yaml
	---
	[--scope=<scope>]
	: Can be all, global, ms_global, blog, or old tables. Defaults to all.
	[--network]
	: List all the tables in a multisite install.
	[--all-tables-with-prefix]
	: List all tables that match the table prefix even if not registered on $wpdb. Overrides --network.
	[--all-tables]
	: List all tables in the database, regardless of the prefix, and even if not registered on $wpdb. Overrides --all-tables-with-prefix.
## EXAMPLES
	    $ wp db size
	    +-------------------+------+
	    | Name              | Size |
	    +-------------------+------+
	    | wordpress_default | 6 MB |
	    +-------------------+------+
	    $ wp db size --tables
	    +-----------------------+-------+
	    | Name                  | Size  |
	    +-----------------------+-------+
	    | wp_users              | 64 KB |
	    | wp_usermeta           | 48 KB |
	    | wp_posts              | 80 KB |
	    | wp_comments           | 96 KB |
	    | wp_links              | 32 KB |
	    | wp_options            | 32 KB |
	    | wp_postmeta           | 48 KB |
	    | wp_terms              | 48 KB |
	    | wp_term_taxonomy      | 48 KB |
	    | wp_term_relationships | 32 KB |
	    | wp_termmeta           | 48 KB |
	    | wp_commentmeta        | 48 KB |
	    +-----------------------+-------+
	    $ wp db size --size_format=b
	    5865472
	    $ wp db size --size_format=kb
	    5728
	    $ wp db size --size_format=mb
	    6
	
 */
package size
import utils "github.com/bukowa/gowpcli"

// Size //Displays the database name and size.
type Size struct {
    SizeFormat string // [--size_format=<format>]
    Tables bool // [--tables]
    HumanReadable bool // [--human-readable]
    Format string // [--format=<format>]
    Scope string // [--scope=<scope>]
    Network bool // [--network]
    AllTablesWithPrefix bool // [--all-tables-with-prefix]
    AllTables bool // [--all-tables]
}

func (s Size) Args() []string {
    var args = []string{"db", "size"}
    args = utils.MakeArg(args, "[--size_format=<format>]", s.SizeFormat)
    args = utils.MakeArg(args, "[--tables]", s.Tables)
    args = utils.MakeArg(args, "[--human-readable]", s.HumanReadable)
    args = utils.MakeArg(args, "[--format=<format>]", s.Format)
    args = utils.MakeArg(args, "[--scope=<scope>]", s.Scope)
    args = utils.MakeArg(args, "[--network]", s.Network)
    args = utils.MakeArg(args, "[--all-tables-with-prefix]", s.AllTablesWithPrefix)
    args = utils.MakeArg(args, "[--all-tables]", s.AllTables)
    return args
}

