/*
## OPTIONS
	<key>
	: Key for the alias.
	[--config=<config>]
	: Config file to be considered for operations.
	---
	options:
	  - global
	  - project
	---
## EXAMPLES
	    # Delete alias.
	    $ wp cli alias delete @prod
	    Success: Deleted '@prod' alias.
	    # Delete project alias.
	    $ wp cli alias delete @prod --config=project
	    Success: Deleted '@prod' alias.
	
 */
package delete
import utils "github.com/bukowa/gowpcli"

// Delete //Deletes an alias.
type Delete struct {
    Key string // <key>
    Config string // [--config=<config>]
}

func (d Delete) Args() []string {
    var args = []string{"cli", "alias", "delete"}
    args = utils.MakeArg(args, "<key>", d.Key)
    args = utils.MakeArg(args, "[--config=<config>]", d.Config)
    return args
}

