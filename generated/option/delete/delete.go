/*
## OPTIONS
	<key>...
	: Key for the option.
## EXAMPLES
	    # Delete an option.
	    $ wp option delete my_option
	    Success: Deleted 'my_option' option.
	    # Delete multiple options.
	    $ wp option delete option_one option_two option_three
	    Success: Deleted 'option_one' option.
	    Success: Deleted 'option_two' option.
	    Warning: Could not delete 'option_three' option. Does it exist?
	
 */
package delete
import utils "github.com/bukowa/gowpcli"

// Delete //Deletes an option.
type Delete struct {
    Key []string // <key>...
}

func (d Delete) Args() []string {
    var args = []string{"option", "delete"}
    args = utils.MakeArg(args, "<key>...", d.Key)
    return args
}

