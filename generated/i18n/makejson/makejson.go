/*
## INFO
	For JavaScript internationalization purposes, WordPress requires translations to be split up into
	one Jed-formatted JSON file per JavaScript source file.
	See https://make.wordpress.org/core/2018/11/09/new-javascript-i18n-support-in-wordpress/ to learn more
	about WordPress JavaScript internationalization.
## OPTIONS
	<source>
	: Path to an existing PO file or a directory containing multiple PO files.
	[<destination>]
	: Path to the destination directory for the resulting JSON files. Defaults to the source directory.
	[--purge]
	: Whether to purge the strings that were extracted from the original source file. Defaults to true, use `--no-purge` to skip the removal.
	[--pretty-print]
	: Pretty-print resulting JSON files.
## EXAMPLES
	    # Create JSON files for all PO files in the languages directory
	    $ wp i18n make-json languages
	    # Create JSON files for my-plugin-de_DE.po and leave the PO file untouched.
	    $ wp i18n make-json my-plugin-de_DE.po /tmp --no-purge
	
 */
package makejson
import utils "github.com/bukowa/gowpcli"

// MakeJson //Extract JavaScript strings from PO files and add them to individual JSON files.
type MakeJson struct {
    Source string // <source>
    Destination string // [<destination>]
    Purge bool // [--purge]
    PrettyPrint bool // [--pretty-print]
}

func (m MakeJson) Args() []string {
    var args = []string{"i18n", "make-json"}
    args = utils.MakeArg(args, "<source>", m.Source)
    args = utils.MakeArg(args, "[<destination>]", m.Destination)
    args = utils.MakeArg(args, "[--purge]", m.Purge)
    args = utils.MakeArg(args, "[--pretty-print]", m.PrettyPrint)
    return args
}

