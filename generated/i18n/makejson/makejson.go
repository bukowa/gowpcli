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
	[--update-mo-files]
	: Whether MO files should be updated as well after updating PO files.
	Only has an effect when used in combination with `--purge`.
	[--pretty-print]
	: Pretty-print resulting JSON files.
	[--use-map=<paths_or_maps>]
	: Whether to use a mapping file for the strings, as a JSON value, array to specify multiple.
	Each element can either be a string (file path) or object (map).
## EXAMPLES
	    # Create JSON files for all PO files in the languages directory
	    $ wp i18n make-json languages
	    # Create JSON files for my-plugin-de_DE.po and leave the PO file untouched.
	    $ wp i18n make-json my-plugin-de_DE.po /tmp --no-purge
	    # Create JSON files with mapping
	    $ wp i18n make-json languages --use-map=build/map.json
	    # Create JSON files with multiple mappings
	    $ wp i18n make-json languages '--use-map=["build/map.json","build/map2.json"]'
	    # Create JSON files with object mapping
	    $ wp i18n make-json languages '--use-map={"source/index.js":"build/index.js"}'
	
 */
package makejson
import utils "github.com/bukowa/gowpcli"

// MakeJson //Extract JavaScript strings from PO files and add them to individual JSON files.
type MakeJson struct {
    Source string // <source>
    Destination string // [<destination>]
    Purge bool // [--purge]
    UpdateMoFiles bool // [--update-mo-files]
    PrettyPrint bool // [--pretty-print]
    UseMap string // [--use-map=<paths_or_maps>]
}

func (m MakeJson) Args() []string {
    var args = []string{"i18n", "make-json"}
    args = utils.MakeArg(args, "<source>", m.Source)
    args = utils.MakeArg(args, "[<destination>]", m.Destination)
    args = utils.MakeArg(args, "[--purge]", m.Purge)
    args = utils.MakeArg(args, "[--update-mo-files]", m.UpdateMoFiles)
    args = utils.MakeArg(args, "[--pretty-print]", m.PrettyPrint)
    args = utils.MakeArg(args, "[--use-map=<paths_or_maps>]", m.UseMap)
    return args
}

