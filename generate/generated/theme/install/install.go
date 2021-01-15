/*
## OPTIONS
	<theme|zip|url>...
	: One or more themes to install. Accepts a theme slug, the path to a local zip file, or a URL to a remote zip file.
	[--version=<version>]
	: If set, get that particular version from wordpress.org, instead of the
	stable version.
	[--force]
	: If set, the command will overwrite any installed version of the theme, without prompting
	for confirmation.
	[--activate]
	: If set, the theme will be activated immediately after install.
## EXAMPLES
	    # Install the latest version from wordpress.org and activate
	    $ wp theme install twentysixteen --activate
	    Installing Twenty Sixteen (1.2)
	    Downloading install package from http://downloads.wordpress.org/theme/twentysixteen.1.2.zip...
	    Unpacking the package...
	    Installing the theme...
	    Theme installed successfully.
	    Activating 'twentysixteen'...
	    Success: Switched to 'Twenty Sixteen' theme.
	    # Install from a local zip file
	    $ wp theme install ../my-theme.zip
	    # Install from a remote zip file
	    $ wp theme install http://s3.amazonaws.com/bucketname/my-theme.zip?AWSAccessKeyId=123&Expires=456&Signature=abcdef
	
 */
package install
import utils "github.com/bukowa/gowpcli"

// Install //Installs one or more themes.
type Install struct {
    Themezipurl []string // <theme|zip|url>...
    Version string // [--version=<version>]
    Force bool // [--force]
    Activate bool // [--activate]
}

func (i Install) Args() []string {
    var args = []string{"theme", "install"}
    args = utils.MakeArg(args, "<theme|zip|url>...", i.Themezipurl)
    args = utils.MakeArg(args, "[--version=<version>]", i.Version)
    args = utils.MakeArg(args, "[--force]", i.Force)
    args = utils.MakeArg(args, "[--activate]", i.Activate)
    return args
}

