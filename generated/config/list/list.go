/*
## OPTIONS
	[<filter>...]
	: Name or partial name to filter the list by.
	[--fields=<fields>]
	: Limit the output to specific fields. Defaults to all fields.
	[--format=<format>]
	: Render output in a particular format.
	Dotenv is limited to non-object values.
	---
	default: table
	options:
	  - table
	  - csv
	  - json
	  - yaml
	  - dotenv
	---
	[--strict]
	: Enforce strict matching when a filter is provided.
	[--config-file=<path>]
	: Specify the file path to the config file to be read. Defaults to the root of the
	WordPress installation and the filename "wp-config.php".
## EXAMPLES
	    # List constants and variables defined in wp-config.php file.
	    $ wp config list
	    +------------------+------------------------------------------------------------------+----------+
	    | key              | value                                                            | type     |
	    +------------------+------------------------------------------------------------------+----------+
	    | table_prefix     | wp_                                                              | variable |
	    | DB_NAME          | wp_cli_test                                                      | constant |
	    | DB_USER          | root                                                             | constant |
	    | DB_PASSWORD      | root                                                             | constant |
	    | AUTH_KEY         | r6+@shP1yO&$)1gdu.hl[/j;7Zrvmt~o;#WxSsa0mlQOi24j2cR,7i+QM/#7S:o^ | constant |
	    | SECURE_AUTH_KEY  | iO-z!_m--YH$Tx2tf/&V,YW*13Z_HiRLqi)d?$o-tMdY+82pK$`T.NYW~iTLW;xp | constant |
	    +------------------+------------------------------------------------------------------+----------+
	    # List only database user and password from wp-config.php file.
	    $ wp config list DB_USER DB_PASSWORD --strict
	    +------------------+-------+----------+
	    | key              | value | type     |
	    +------------------+-------+----------+
	    | DB_USER          | root  | constant |
	    | DB_PASSWORD      | root  | constant |
	    +------------------+-------+----------+
	    # List all salts from wp-config.php file.
	    $ wp config list _SALT
	    +------------------+------------------------------------------------------------------+----------+
	    | key              | value                                                            | type     |
	    +------------------+------------------------------------------------------------------+----------+
	    | AUTH_SALT        | n:]Xditk+_7>Qi=>BmtZHiH-6/Ecrvl(V5ceeGP:{>?;BT^=[B3-0>,~F5z$(+Q$ | constant |
	    | SECURE_AUTH_SALT | ?Z/p|XhDw3w}?c.z%|+BAr|(Iv*H%%U+Du&kKR y?cJOYyRVRBeB[2zF-`(>+LCC | constant |
	    | LOGGED_IN_SALT   | +$@(1{b~Z~s}Cs>8Y]6[m6~TnoCDpE>O%e75u}&6kUH!>q:7uM4lxbB6[1pa_X,q | constant |
	    | NONCE_SALT       | _x+F li|QL?0OSQns1_JZ{|Ix3Jleox-71km/gifnyz8kmo=w-;@AE8W,(fP<N}2 | constant |
	    +------------------+------------------------------------------------------------------+----------+
	
 */
package list
import utils "github.com/bukowa/gowpcli"

// List //Lists variables, constants, and file includes defined in wp-config.php file.
type List struct {
    Filter []string // [<filter>...]
    Fields string // [--fields=<fields>]
    Format string // [--format=<format>]
    Strict bool // [--strict]
    ConfigFile string // [--config-file=<path>]
}

func (l List) Args() []string {
    var args = []string{"config", "list"}
    args = utils.MakeArg(args, "[<filter>...]", l.Filter)
    args = utils.MakeArg(args, "[--fields=<fields>]", l.Fields)
    args = utils.MakeArg(args, "[--format=<format>]", l.Format)
    args = utils.MakeArg(args, "[--strict]", l.Strict)
    args = utils.MakeArg(args, "[--config-file=<path>]", l.ConfigFile)
    return args
}

