/*
## OPTIONS
	<id>
	: The ID of the object.
	[<key>]
	: The name of the meta field to delete.
	[<value>]
	: The value to delete. If omitted, all rows with key will deleted.
	[--all]
	: Delete all meta for the object.
	
 */
package delete
import utils "github.com/bukowa/gowpcli"

// Delete //Delete a meta field.
type Delete struct {
    Id string // <id>
    Key string // [<key>]
    Value string // [<value>]
    All bool // [--all]
}

func (d Delete) Args() []string {
    var args = []string{"comment", "meta", "delete"}
    args = utils.MakeArg(args, "<id>", d.Id)
    args = utils.MakeArg(args, "[<key>]", d.Key)
    args = utils.MakeArg(args, "[<value>]", d.Value)
    args = utils.MakeArg(args, "[--all]", d.All)
    return args
}

