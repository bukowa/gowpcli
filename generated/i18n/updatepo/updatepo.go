/*
## INFO
	This behaves similarly to the [msgmerge](https://www.gnu.org/software/gettext/manual/html_node/msgmerge-Invocation.html) command.
## OPTIONS
	<source>
	: Path to an existing POT file to use for updating
	[<destination>]
	: PO file to update or a directory containing multiple PO files.
	  Defaults to all PO files in the source directory.
	
 */
package updatepo
import utils "github.com/bukowa/gowpcli"

// UpdatePo //Update PO files from a POT file.
type UpdatePo struct {
    Source string // <source>
    Destination string // [<destination>]
}

func (u UpdatePo) Args() []string {
    var args = []string{"i18n", "update-po"}
    args = utils.MakeArg(args, "<source>", u.Source)
    args = utils.MakeArg(args, "[<destination>]", u.Destination)
    return args
}

