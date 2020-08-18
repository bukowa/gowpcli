/*
## INFO
	Packages are required to be a valid Composer package, and can be
	specified as:
	* Package name from WP-CLI's package index.
	* Git URL accessible by the current shell user.
	* Path to a directory on the local machine.
	* Local or remote .zip file.
	Packages are installed to `~/.wp-cli/packages/` by default. Use the
	`WP_CLI_PACKAGES_DIR` environment variable to provide a custom path.
	When installing a local directory, WP-CLI simply registers a
	reference to the directory. If you move or delete the directory, WP-CLI's
	reference breaks.
	When installing a .zip file, WP-CLI extracts the package to
	`~/.wp-cli/packages/local/<package-name>`.
## OPTIONS
	<name|git|path|zip>
	: Name, git URL, directory path, or .zip file for the package to install.
	Names can optionally include a version constraint
	(e.g. wp-cli/server-command:@stable).
## EXAMPLES
	    # Install the latest development version from the package index.
	    $ wp package install wp-cli/server-command
	    Installing package wp-cli/server-command (dev-master)
	    Updating /home/person/.wp-cli/packages/composer.json to require the package...
	    Using Composer to install the package...
	    ---
	    Loading composer repositories with package information
	    Updating dependencies
	    Resolving dependencies through SAT
	    Dependency resolution completed in 0.005 seconds
	    Analyzed 732 packages to resolve dependencies
	    Analyzed 1034 rules to resolve dependencies
	     - Installing package
	    Writing lock file
	    Generating autoload files
	    ---
	    Success: Package installed.
	    # Install the latest stable version.
	    $ wp package install wp-cli/server-command:@stable
	    # Install a package hosted at a git URL.
	    $ wp package install git@github.com:runcommand/hook.git
	    # Install a package in a .zip file.
	    $ wp package install google-sitemap-generator-cli.zip
	
 */
package install
import utils "github.com/bukowa/gowpcli"

// Install //Installs a WP-CLI package.
type Install struct {
    Namegitpathzip string // <name|git|path|zip>
}

func (i Install) Args() []string {
    var args = []string{"package", "install"}
    args = utils.MakeArg(args, "<name|git|path|zip>", i.Namegitpathzip)
    return args
}

