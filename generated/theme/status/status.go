/*
## OPTIONS
	[<theme>]
	: A particular theme to show the status for.
## EXAMPLES
	    $ wp theme status twentysixteen
	    Theme twentysixteen details:
	         Name: Twenty Sixteen
	         Status: Inactive
	         Version: 1.2
	         Author: the WordPress team
	
 */
package status
import utils "github.com/bukowa/gowpcli"

// Status //Reveals the status of one or all themes.
type Status struct {
    Theme string // [<theme>]
}

func (s Status) Args() []string {
    var args = []string{"theme", "status"}
    args = utils.MakeArg(args, "[<theme>]", s.Theme)
    return args
}

