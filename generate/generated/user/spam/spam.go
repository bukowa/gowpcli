/*
## OPTIONS
	<id>...
	: One or more IDs of users to mark as spam.
## EXAMPLES
	    $ wp user spam 123
	    User 123 marked as spam.
	    Success: Spamed 1 of 1 users.
	
 */
package spam
import utils "github.com/bukowa/gowpcli"

// Spam //Marks one or more users as spam.
type Spam struct {
    Id []string // <id>...
}

func (s Spam) Args() []string {
    var args = []string{"user", "spam"}
    args = utils.MakeArg(args, "<id>...", s.Id)
    return args
}

