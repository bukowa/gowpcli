/*
## OPTIONS
	<source>
	: Path to an existing PO file or a directory containing multiple PO files.
	[<destination>]
	: Path to the destination directory for the resulting MO files. Defaults to the source directory.
## EXAMPLES
	    # Create MO files for all PO files in the current directory.
	    $ wp i18n make-mo .
	    # Create a MO file from a single PO file in a specific directory.
	    $ wp i18n make-mo example-plugin-de_DE.po languages
	
 */
package makemo
import utils "github.com/bukowa/gowpcli"

// MakeMo //Create MO files from PO files.
type MakeMo struct {
    Source string // <source>
    Destination string // [<destination>]
}

func (m MakeMo) Args() []string {
    var args = []string{"i18n", "make-mo"}
    args = utils.MakeArg(args, "<source>", m.Source)
    args = utils.MakeArg(args, "[<destination>]", m.Destination)
    return args
}

