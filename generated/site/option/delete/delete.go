/*
## OPTIONS
	<key>
	: Key for the site option.
## EXAMPLES
	    $ wp site option delete my_option
	    Success: Deleted 'my_option' site option.
	
 */
package delete
import utils "github.com/bukowa/gowpcli"

// Delete //Deletes a site option.
type Delete struct {
    Key string // <key>
}

func (d Delete) Args() []string {
    var args = []string{"site", "option", "delete"}
    args = utils.MakeArg(args, "<key>", d.Key)
    return args
}

