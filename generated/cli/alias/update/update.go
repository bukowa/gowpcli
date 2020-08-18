/*
## OPTIONS
	<key>
	: Key for the alias.
	[--set-user=<user>]
	: Set user for alias.
	[--set-url=<url>]
	: Set url for alias.
	[--set-path=<path>]
	: Set path for alias.
	[--set-ssh=<ssh>]
	: Set ssh for alias.
	[--set-http=<http>]
	: Set http for alias.
	[--grouping=<grouping>]
	: For grouping multiple aliases.
	[--config=<config>]
	: Config file to be considered for operations.
	---
	options:
	  - global
	  - project
	---
## EXAMPLES
	    # Update alias.
	    $ wp cli alias update @prod --set-user=newuser --set-path=/new/path/to/wordpress/install/
	    Success: Updated 'prod' alias.
	    # Update project alias.
	    $ wp cli alias update @prod --set-user=newuser --set-path=/new/path/to/wordpress/install/ --config=project
	    Success: Updated 'prod' alias.
	
 */
package update
import utils "github.com/bukowa/gowpcli"

// Update //Updates an alias.
type Update struct {
    Key string // <key>
    SetUser string // [--set-user=<user>]
    SetUrl string // [--set-url=<url>]
    SetPath string // [--set-path=<path>]
    SetSsh string // [--set-ssh=<ssh>]
    SetHttp string // [--set-http=<http>]
    Grouping string // [--grouping=<grouping>]
    Config string // [--config=<config>]
}

func (u Update) Args() []string {
    var args = []string{"cli", "alias", "update"}
    args = utils.MakeArg(args, "<key>", u.Key)
    args = utils.MakeArg(args, "[--set-user=<user>]", u.SetUser)
    args = utils.MakeArg(args, "[--set-url=<url>]", u.SetUrl)
    args = utils.MakeArg(args, "[--set-path=<path>]", u.SetPath)
    args = utils.MakeArg(args, "[--set-ssh=<ssh>]", u.SetSsh)
    args = utils.MakeArg(args, "[--set-http=<http>]", u.SetHttp)
    args = utils.MakeArg(args, "[--grouping=<grouping>]", u.Grouping)
    args = utils.MakeArg(args, "[--config=<config>]", u.Config)
    return args
}

