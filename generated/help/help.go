/*
## OPTIONS
	[<command>...]
	: Get help on a specific command.
## EXAMPLES
	    # get help for `core` command
	    wp help core
	    # get help for `core download` subcommand
	    wp help core download
	
 */
package help
import utils "github.com/bukowa/gowpcli"

// Help //Gets help on WP-CLI, or on a specific command.
type Help struct {
    Command []string // [<command>...]
}

func (h Help) Args() []string {
    var args = []string{"help"}
    args = utils.MakeArg(args, "[<command>...]", h.Command)
    return args
}

