/*
## OPTIONS
	<user>
	: The user login, user email, or user ID of the user to update the application password for.
	<uuid>
	: The universally unique ID of the application password.
	[--<field>=<value>]
	: Update the <field> with a new <value>. Currently supported fields: name.
## EXAMPLES
	    # Update an existing application password
	    $ wp user application-password update 123 6633824d-c1d7-4f79-9dd5-4586f734d69e --name=newappname
	    Success: Updated application password.
	
 */
package update
import utils "github.com/bukowa/gowpcli"

// Update //Updates an existing application password.
type Update struct {
    User string // <user>
    Uuid string // <uuid>
    FieldMap map[string]string // [--<field>=<value>]
}

func (u Update) Args() []string {
    var args = []string{"user", "application-password", "update"}
    args = utils.MakeArg(args, "<user>", u.User)
    args = utils.MakeArg(args, "<uuid>", u.Uuid)
    args = utils.MakeArg(args, "[--<field>=<value>]", u.FieldMap)
    return args
}

