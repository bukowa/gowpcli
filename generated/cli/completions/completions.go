/*
## OPTIONS
	--line=<line>
	: The current command line to be executed.
	--point=<point>
	: The index to the current cursor position relative to the beginning of the command.
## EXAMPLES
	    # Generate tab completion strings.
	    $ wp cli completions --line='wp eva' --point=100
	    eval
	    eval-file
	
 */
package completions
import utils "github.com/bukowa/gowpcli"

// Completions //Generates tab completion strings.
type Completions struct {
    Line string // --line=<line>
    Point string // --point=<point>
}

func (c Completions) Args() []string {
    var args = []string{"cli", "completions"}
    args = utils.MakeArg(args, "--line=<line>", c.Line)
    args = utils.MakeArg(args, "--point=<point>", c.Point)
    return args
}

