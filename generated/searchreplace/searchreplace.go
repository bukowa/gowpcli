/*
## INFO
	Searches through all rows in a selection of tables and replaces
	appearances of the first string with the second string.
	By default, the command uses tables registered to the `$wpdb` object. On
	multisite, this will just be the tables for the current site unless
	`--network` is specified.
	Search/replace intelligently handles PHP serialized data, and does not
	change primary key values.
## OPTIONS
	<old>
	: A string to search for within the database.
	<new>
	: Replace instances of the first string with this new string.
	[<table>...]
	: List of database tables to restrict the replacement to. Wildcards are
	supported, e.g. `'wp_*options'` or `'wp_post*'`.
	[--dry-run]
	: Run the entire search/replace operation and show report, but don't save
	changes to the database.
	[--network]
	: Search/replace through all the tables registered to $wpdb in a
	multisite install.
	[--all-tables-with-prefix]
	: Enable replacement on any tables that match the table prefix even if
	not registered on $wpdb.
	[--all-tables]
	: Enable replacement on ALL tables in the database, regardless of the
	prefix, and even if not registered on $wpdb. Overrides --network
	and --all-tables-with-prefix.
	[--export[=<file>]]
	: Write transformed data as SQL file instead of saving replacements to
	the database. If <file> is not supplied, will output to STDOUT.
	[--export_insert_size=<rows>]
	: Define number of rows in single INSERT statement when doing SQL export.
	You might want to change this depending on your database configuration
	(e.g. if you need to do fewer queries). Default: 50
	[--skip-tables=<tables>]
	: Do not perform the replacement on specific tables. Use commas to
	specify multiple tables. Wildcards are supported, e.g. `'wp_*options'` or `'wp_post*'`.
	[--skip-columns=<columns>]
	: Do not perform the replacement on specific columns. Use commas to
	specify multiple columns.
	[--include-columns=<columns>]
	: Perform the replacement on specific columns. Use commas to
	specify multiple columns.
	[--precise]
	: Force the use of PHP (instead of SQL) which is more thorough,
	but slower.
	[--recurse-objects]
	: Enable recursing into objects to replace strings. Defaults to true;
	pass --no-recurse-objects to disable.
	[--verbose]
	: Prints rows to the console as they're updated.
	[--regex]
	: Runs the search using a regular expression (without delimiters).
	Warning: search-replace will take about 15-20x longer when using --regex.
	[--regex-flags=<regex-flags>]
	: Pass PCRE modifiers to regex search-replace (e.g. 'i' for case-insensitivity).
	[--regex-delimiter=<regex-delimiter>]
	: The delimiter to use for the regex. It must be escaped if it appears in the search string. The default value is the result of `chr(1)`.
	[--regex-limit=<regex-limit>]
	: The maximum possible replacements for the regex per row (or per unserialized data bit per row). Defaults to -1 (no limit).
	[--format=<format>]
	: Render output in a particular format.
	---
	default: table
	options:
	  - table
	  - count
	---
	[--report]
	: Produce report. Defaults to true.
	[--report-changed-only]
	: Report changed fields only. Defaults to false, unless logging, when it defaults to true.
	[--log[=<file>]]
	: Log the items changed. If <file> is not supplied or is "-", will output to STDOUT.
	Warning: causes a significant slow down, similar or worse to enabling --precise or --regex.
	[--before_context=<num>]
	: For logging, number of characters to display before the old match and the new replacement. Default 40. Ignored if not logging.
	[--after_context=<num>]
	: For logging, number of characters to display after the old match and the new replacement. Default 40. Ignored if not logging.
## EXAMPLES
	    # Search and replace but skip one column
	    $ wp search-replace 'http://example.test' 'http://example.com' --skip-columns=guid
	    # Run search/replace operation but dont save in database
	    $ wp search-replace 'foo' 'bar' wp_posts wp_postmeta wp_terms --dry-run
	    # Run case-insensitive regex search/replace operation (slow)
	    $ wp search-replace '\[foo id="([0-9]+)"' '[bar id="\1"' --regex --regex-flags='i'
	    # Turn your production multisite database into a local dev database
	    $ wp search-replace --url=example.com example.com example.test 'wp_*options' wp_blogs
	    # Search/replace to a SQL file without transforming the database
	    $ wp search-replace foo bar --export=database.sql
	    # Bash script: Search/replace production to development url (multisite compatible)
	    #!/bin/bash
	    if $(wp --url=http://example.com core is-installed --network); then
	        wp search-replace --url=http://example.com 'http://example.com' 'http://example.test' --recurse-objects --network --skip-columns=guid --skip-tables=wp_users
	    else
	        wp search-replace 'http://example.com' 'http://example.test' --recurse-objects --skip-columns=guid --skip-tables=wp_users
	    fi
	
 */
package searchreplace
import utils "github.com/bukowa/gowpcli"

// SearchReplace //Searches/replaces strings in the database.
type SearchReplace struct {
    Old string // <old>
    New string // <new>
    Table []string // [<table>...]
    DryRun bool // [--dry-run]
    Network bool // [--network]
    AllTablesWithPrefix bool // [--all-tables-with-prefix]
    AllTables bool // [--all-tables]
    Export string // [--export[=<file>]]
    ExportInsertSize string // [--export_insert_size=<rows>]
    SkipTables string // [--skip-tables=<tables>]
    SkipColumns string // [--skip-columns=<columns>]
    IncludeColumns string // [--include-columns=<columns>]
    Precise bool // [--precise]
    RecurseObjects bool // [--recurse-objects]
    Verbose bool // [--verbose]
    Regex bool // [--regex]
    RegexFlags string // [--regex-flags=<regex-flags>]
    RegexDelimiter string // [--regex-delimiter=<regex-delimiter>]
    RegexLimit string // [--regex-limit=<regex-limit>]
    Format string // [--format=<format>]
    Report bool // [--report]
    ReportChangedOnly bool // [--report-changed-only]
    Log string // [--log[=<file>]]
    BeforeContext string // [--before_context=<num>]
    AfterContext string // [--after_context=<num>]
}

func (s SearchReplace) Args() []string {
    var args = []string{"search-replace"}
    args = utils.MakeArg(args, "<old>", s.Old)
    args = utils.MakeArg(args, "<new>", s.New)
    args = utils.MakeArg(args, "[<table>...]", s.Table)
    args = utils.MakeArg(args, "[--dry-run]", s.DryRun)
    args = utils.MakeArg(args, "[--network]", s.Network)
    args = utils.MakeArg(args, "[--all-tables-with-prefix]", s.AllTablesWithPrefix)
    args = utils.MakeArg(args, "[--all-tables]", s.AllTables)
    args = utils.MakeArg(args, "[--export[=<file>]]", s.Export)
    args = utils.MakeArg(args, "[--export_insert_size=<rows>]", s.ExportInsertSize)
    args = utils.MakeArg(args, "[--skip-tables=<tables>]", s.SkipTables)
    args = utils.MakeArg(args, "[--skip-columns=<columns>]", s.SkipColumns)
    args = utils.MakeArg(args, "[--include-columns=<columns>]", s.IncludeColumns)
    args = utils.MakeArg(args, "[--precise]", s.Precise)
    args = utils.MakeArg(args, "[--recurse-objects]", s.RecurseObjects)
    args = utils.MakeArg(args, "[--verbose]", s.Verbose)
    args = utils.MakeArg(args, "[--regex]", s.Regex)
    args = utils.MakeArg(args, "[--regex-flags=<regex-flags>]", s.RegexFlags)
    args = utils.MakeArg(args, "[--regex-delimiter=<regex-delimiter>]", s.RegexDelimiter)
    args = utils.MakeArg(args, "[--regex-limit=<regex-limit>]", s.RegexLimit)
    args = utils.MakeArg(args, "[--format=<format>]", s.Format)
    args = utils.MakeArg(args, "[--report]", s.Report)
    args = utils.MakeArg(args, "[--report-changed-only]", s.ReportChangedOnly)
    args = utils.MakeArg(args, "[--log[=<file>]]", s.Log)
    args = utils.MakeArg(args, "[--before_context=<num>]", s.BeforeContext)
    args = utils.MakeArg(args, "[--after_context=<num>]", s.AfterContext)
    return args
}

