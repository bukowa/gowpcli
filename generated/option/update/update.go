/*
## OPTIONS
	<key>
	: The name of the option to update.
	[<value>]
	: The new value. If ommited, the value is read from STDIN.
	[--autoload=<autoload>]
	: Requires WP 4.2. Should this option be automatically loaded.
	---
	options:
	  - 'yes'
	  - 'no'
	---
	[--format=<format>]
	: The serialization format for the value.
	---
	default: plaintext
	options:
	  - plaintext
	  - json
	---
## EXAMPLES
	    # Update an option by reading from a file.
	    $ wp option update my_option < value.txt
	    Success: Updated 'my_option' option.
	    # Update one option on multiple sites using xargs.
	    $ wp site list --field=url | xargs -n1 -I {} sh -c 'wp --url={} option update my_option my_value'
	    Success: Updated 'my_option' option.
	    Success: Updated 'my_option' option.
	    # Update site blog name.
	    $ wp option update blogname "Random blog name"
	    Success: Updated 'blogname' option.
	    # Update site blog description.
	    $ wp option update blogdescription "Some random blog description"
	    Success: Updated 'blogdescription' option.
	    # Update admin email address.
	    $ wp option update admin_email someone@example.com
	    Success: Updated 'admin_email' option.
	    # Set the default role.
	    $ wp option update default_role author
	    Success: Updated 'default_role' option.
	    # Set the timezone string.
	    $ wp option update timezone_string "America/New_York"
	    Success: Updated 'timezone_string' option.
	
 */
package update
import utils "github.com/bukowa/gowpcli"

// Update //Updates an option value.
type Update struct {
    Key string // <key>
    Value string // [<value>]
    Autoload string // [--autoload=<autoload>]
    Format string // [--format=<format>]
}

func (u Update) Args() []string {
    var args = []string{"option", "update"}
    args = utils.MakeArg(args, "<key>", u.Key)
    args = utils.MakeArg(args, "[<value>]", u.Value)
    args = utils.MakeArg(args, "[--autoload=<autoload>]", u.Autoload)
    args = utils.MakeArg(args, "[--format=<format>]", u.Format)
    return args
}

