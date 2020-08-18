/*
## INFO
	Downloads the language pack from WordPress.org.
## OPTIONS
	[<plugin>]
	: Plugin to install language for.
	[--all]
	: If set, languages for all plugins will be installed.
	<language>...
	: Language code to install.
	[--format=<format>]
	: Render output in a particular format. Used when installing languages for all plugins.
	---
	default: table
	options:
	  - table
	  - csv
	  - json
	  - summary
	---
## EXAMPLES
	    # Install the Japanese language for Akismet.
	    $ wp language plugin install akismet ja
	    Downloading translation from https://downloads.wordpress.org/translation/plugin/akismet/4.0.3/ja.zip...
	    Unpacking the update...
	    Installing the latest version...
	    Translation updated successfully.
	    Language 'ja' installed.
	    Success: Installed 1 of 1 languages.
	
 */
package install
import utils "github.com/bukowa/gowpcli"

// Install //Installs a given language for a plugin.
type Install struct {
    Plugin string // [<plugin>]
    All bool // [--all]
    Language []string // <language>...
    Format string // [--format=<format>]
}

func (i Install) Args() []string {
    var args = []string{"language", "plugin", "install"}
    args = utils.MakeArg(args, "[<plugin>]", i.Plugin)
    args = utils.MakeArg(args, "[--all]", i.All)
    args = utils.MakeArg(args, "<language>...", i.Language)
    args = utils.MakeArg(args, "[--format=<format>]", i.Format)
    return args
}

