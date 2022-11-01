/*
## INFO
	Returns exit code 0 when active, 1 when not active.
## OPTIONS
	<plugin>
	: The plugin to check.
	[--network]
	: If set, check if plugin is network-activated.
## EXAMPLES
	    # Check whether plugin is Active; exit status 0 if active, otherwise 1
	    $ wp plugin is-active hello
	    $ echo $?
	    1
	
 */
package isactive
import utils "github.com/bukowa/gowpcli"

// IsActive //Checks if a given plugin is active.
type IsActive struct {
    Plugin string // <plugin>
    Network bool // [--network]
}

func (i IsActive) Args() []string {
    var args = []string{"plugin", "is-active"}
    args = utils.MakeArg(args, "<plugin>", i.Plugin)
    args = utils.MakeArg(args, "[--network]", i.Network)
    return args
}

