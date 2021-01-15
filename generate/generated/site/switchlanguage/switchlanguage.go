/*
## OPTIONS
	<language>
	: Language code to activate.
## EXAMPLES
	    $ wp site switch-language ja
	    Success: Language activated.
	
 */
package switchlanguage
import utils "github.com/bukowa/gowpcli"

// SwitchLanguage //Activates a given language.
type SwitchLanguage struct {
    Language string // <language>
}

func (s SwitchLanguage) Args() []string {
    var args = []string{"site", "switch-language"}
    args = utils.MakeArg(args, "<language>", s.Language)
    return args
}

