/*
## OPTIONS
	<language>
	: Language code to activate.
## EXAMPLES
	    $ wp language core activate ja
	    Success: Language activated.
	
 */
package activate
import utils "github.com/bukowa/gowpcli"

// Activate //Activates a given language.
type Activate struct {
    Language string // <language>
}

func (a Activate) Args() []string {
    var args = []string{"language", "core", "activate"}
    args = utils.MakeArg(args, "<language>", a.Language)
    return args
}

