/*
## INFO
	Downloads and extracts WordPress core files to the specified path. Uses
	current directory when no path is specified. Downloaded build is verified
	to have the correct md5 and then cached to the local filesytem.
	Subsequent uses of command will use the local cache if it still exists.
## OPTIONS
	[<download-url>]
	: Download directly from a provided URL instead of fetching the URL from the wordpress.org servers.
	[--path=<path>]
	: Specify the path in which to install WordPress. Defaults to current
	directory.
	[--locale=<locale>]
	: Select which language you want to download.
	[--version=<version>]
	: Select which version you want to download. Accepts a version number, 'latest' or 'nightly'.
	[--skip-content]
	: Download WP without the default themes and plugins.
	[--force]
	: Overwrites existing files, if present.
## EXAMPLES
	    $ wp core download --locale=nl_NL
	    Downloading WordPress 4.5.2 (nl_NL)...
	    md5 hash verified: c5366d05b521831dd0b29dfc386e56a5
	    Success: WordPress downloaded.
	
 */
package download
import utils "github.com/bukowa/gowpcli"

// Download //Downloads core WordPress files.
type Download struct {
    DownloadUrl string // [<download-url>]
    Path string // [--path=<path>]
    Locale string // [--locale=<locale>]
    Version string // [--version=<version>]
    SkipContent bool // [--skip-content]
    Force bool // [--force]
}

func (d Download) Args() []string {
    var args = []string{"core", "download"}
    args = utils.MakeArg(args, "[<download-url>]", d.DownloadUrl)
    args = utils.MakeArg(args, "[--path=<path>]", d.Path)
    args = utils.MakeArg(args, "[--locale=<locale>]", d.Locale)
    args = utils.MakeArg(args, "[--version=<version>]", d.Version)
    args = utils.MakeArg(args, "[--skip-content]", d.SkipContent)
    args = utils.MakeArg(args, "[--force]", d.Force)
    return args
}

